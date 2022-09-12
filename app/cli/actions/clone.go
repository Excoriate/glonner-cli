package actions

import (
	"fmt"
	logger "github.com/glonner/pkg/log"
	"github.com/glonner/pkg/system"
	"github.com/glonner/pkg/ux"
	"github.com/glonner/services"
	"github.com/pterm/pterm"
	"strings"
	"time"
)

type CloneActionArgs struct {
	Dir          string
	SkipIfExists bool
	ForceIfExist bool
	StorageCheck bool
	DryRun       bool
}

type CloneRepoData struct {
	Name         string
	Dir          string
	GitURL       string
	IsCloned     bool
	IsSkipped    bool
	IsOverridden bool
	IsDryRun     bool
	IsError      bool
}

type CloneActionResult struct {
	RepositoryName  string
	CloneURL        string
	Dir             string
	DirCloned       string
	SizeOccupied    int64
	SizeCloned      int64
	ReposSkipped    []CloneRepoData
	ReposOverridden []CloneRepoData
	ReposFailed     []CloneRepoData
}

type CloneDirType struct {
	isNew         bool
	isEmpty       bool
	isWithContent bool
	rootDir       string
}

var filesToExclude = []string{".git", ".gitignore", ".gitattributes", ".gitmodules", ".gitkeep",
	".DS_Store"}

func getDirConfig(dir string) CloneDirType {
	// 1. Check whether the directory exists. If not, it'll be marked to be created.
	// 2. Check whether the directory is empty. Only, if exists.
	// 3. Check whether the directory is with content. Only, if exists.
	var dirType CloneDirType
	dirType.rootDir = dir

	if !system.CheckIfDirExists(dir) {
		dirType.isNew = true
		dirType.isEmpty = true
		dirType.isWithContent = false
	} else {
		if system.CheckIfDirHasContent(dir, filesToExclude) {
			dirType.isNew = false
			dirType.isEmpty = false
			dirType.isWithContent = true
		} else {
			dirType.isNew = false
			dirType.isEmpty = true
			dirType.isWithContent = false
		}
	}

	if dirType.isNew {
		ux.OutWarn(fmt.Sprintf("The directory %s does not exist. "+
			"It will be created automatically", dir), "")
	} else {
		if dirType.isEmpty {
			ux.OutInfo(fmt.Sprintf("The directory %s is empty. "+
				"It will be used to clone the repositories", dir), "")
		} else {
			if dirType.isWithContent {
				ux.OutWarn(fmt.Sprintf("The directory %s is not empty (content detected). "+
					"It will be used to clone the repositories", dir), "")
			}
		}
	}

	return dirType
}

func checkIncompatibleArgs(args CloneActionArgs) bool {
	if args.SkipIfExists && args.ForceIfExist {
		return true
	}

	return false
}

func getRepositoriesFromGitHub(logger *logger.ILogger, org string, token string,
	dir string) ([]CloneActionResult,
	error) {
	ghSvc := services.NewGitHubSvc(logger)
	repositories, err := ghSvc.GetRepositories(org, token)

	if err != nil {
		// TODO: Handler error, accordingly.
		return nil, err
	}

	var data []CloneActionResult

	for _, repo := range repositories {
		repoDataSimplified := ghSvc.GetSimplifiedRepositoryData(repo)
		for _, item := range repoDataSimplified {
			data = append(data, CloneActionResult{
				RepositoryName: item.Name,
				CloneURL:       item.CloneURL,
				Dir:            dir,
				DirCloned:      fmt.Sprintf("%s%s", dir, item.Name),
				SizeOccupied:   0,
				SizeCloned:     int64(item.Size),
				// Stats, TODO: Implement this functionality later
				ReposFailed:     []CloneRepoData{},
				ReposSkipped:    []CloneRepoData{},
				ReposOverridden: []CloneRepoData{},
			})
		}
	}

	return data, nil
}

func getRootDirToCloneRepos(dir string) string {
	rootDirPathToCLone := fmt.Sprintf("%s/%s", system.GetHomeDir(), dir)

	if !strings.HasSuffix(rootDirPathToCLone, "/") {
		rootDirPathToCLone = fmt.Sprintf("%s/", rootDirPathToCLone)
	}

	return rootDirPathToCLone
}

func createRootDirIfNotExists(rootDir string) (bool, error) {
	if !system.CheckIfDirExists(rootDir) {
		ux.OutWarn(fmt.Sprintf("The directory %s does not exist. "+
			"It will be created automatically", rootDir), "")

		err := system.CreateDir(rootDir)

		if err != nil {
			return false, err
		}

		return true, nil
	}
	return false, nil
}

func clone(s services.IGitHubSvc, repoDir string, repoName string, gitURL string,
	dryRun bool) CloneRepoData {

	if dryRun {
		return CloneRepoData{
			Name:     repoName,
			Dir:      repoDir,
			GitURL:   gitURL,
			IsCloned: true,
			IsDryRun: true,
		}
	}

	isCloned, err := s.CloneRepository(gitURL, repoDir)

	if err != nil {
		if err != nil {
			ux.OutError(fmt.Sprintf("The repository %s could not be cloned due to an error. "+
				"Error: %s", repoName,
				err.Error()), "", false)
		}

		return CloneRepoData{
			Name:     repoName,
			Dir:      repoDir,
			GitURL:   gitURL,
			IsCloned: false,
			IsError:  true,
			IsDryRun: false,
		}
	}

	if !isCloned {
		if err != nil {
			ux.OutError(fmt.Sprintf("The repository %s could not be cloned.", repoName), "", false)
		}

		return CloneRepoData{
			Name:     repoName,
			Dir:      repoDir,
			GitURL:   gitURL,
			IsCloned: false,
			IsDryRun: false,
			IsError:  false,
		}

	}

	return CloneRepoData{}
}

func setupProgress(data []CloneActionResult, ghOrg string) (*pterm.ProgressbarPrinter,
	*ux.ProgressBar) {

	pbTitle := fmt.Sprintf("Cloning repositories of GitHub organisation: %s", ghOrg)
	pbProgressMsg := "Cloning Git Repository"

	pb := ux.NewProgressBar(pbTitle, 1, pbProgressMsg, "", "", "")

	var pbItems []string
	for _, item := range data {
		pbItems = append(pbItems, item.RepositoryName)
	}

	pbStarted := pb.Start(pbItems)

	return pbStarted, pb
}

func CloneAction(globals GlobalRequiredArgs, args CloneActionArgs, logger *logger.ILogger) {

	spinnerPreChecks := ux.GetSpinner("Checking pre-requisites", 1)

	// 1. Check if the arguments are compatible.
	// -----------------------------------
	if checkIncompatibleArgs(args) {
		ux.OutError("The arguments --skip-if-exists and --force-if-exists are incompatible. "+
			"Please, use only one of them.", "", false)
		return
	}

	rootDir := getRootDirToCloneRepos(args.Dir)

	// 2. Get the directory config.
	// 2.1. Identify what type of directory is: new, empty, with content.
	// 2.2. Warn user if the directory is not empty.
	// -----------------------------------
	dirConfig := getDirConfig(rootDir)

	if dirConfig.isNew {
		isCreated, err := createRootDirIfNotExists(dirConfig.rootDir)

		if err != nil || !isCreated {
			ux.OutError(fmt.Sprintf("The directory %s could not be created.", dirConfig.rootDir),
				"", false)
		} else {
			ux.OutInfo(fmt.Sprintf("The directory %s was created successfully", dirConfig.rootDir), "")
		}
	}

	// 3. Get the repositories from GitHub.
	// -----------------------------------
	ghSvc := services.NewGitHubSvc(logger)
	repositories, err := getRepositoriesFromGitHub(logger, globals.Organization, globals.Token, rootDir)

	if err != nil {
		ux.OutError("Could not fetch repository data.", "", false)
	}

	spinnerPreChecks.Success("Pre-requisites checked successfully")

	// Summary of the results expected after the clone-process finishes.
	// -----------------------------------
	var summaryReposCloned []CloneRepoData
	var summaryReposFailed []CloneRepoData
	var summaryReposSkipped []CloneRepoData
	var summaryReposOverridden []CloneRepoData

	pbInstance, pbLib := setupProgress(repositories, globals.Organization)

	// 4. Clone the repositories.
	// 4.1. Destination directory is empty, or it is fairly new.
	// -----------------------------------
	if dirConfig.isNew || dirConfig.isEmpty {
		// If it's new, it was created already, so become empty.
		// If it's empty, it was created already, so it remains empty.
		for _, repo := range repositories {
			specificRepoDir := fmt.Sprintf("%s%s", dirConfig.rootDir, repo.RepositoryName)

			// Increment the progress bar.
			pbInstance.UpdateTitle(fmt.Sprintf("Cloning Git Repository ("+
				"new/empty dir): %s from address: %s",
				repo.RepositoryName, repo.CloneURL))

			time.Sleep(pbLib.SleepTime)

			if !system.CheckIfDirExists(specificRepoDir) {
				err = system.CreateDir(specificRepoDir)
				if err != nil {
					ux.OutError(fmt.Sprintf("The repository directory %s could not be created. "+
						"Error: %s", specificRepoDir, err.Error()), "", false)

					summaryReposFailed = append(summaryReposFailed, CloneRepoData{
						Name:     repo.RepositoryName,
						Dir:      specificRepoDir,
						GitURL:   repo.CloneURL,
						IsCloned: false,
						IsError:  true,
					})

					pbLib.OnFail("", fmt.Sprintf("Git clone failed for repository %s, "+
						"with address: %s", repo.RepositoryName, repo.CloneURL))
				} else {
					repo.DirCloned = specificRepoDir
					cloneResult := clone(ghSvc, repo.DirCloned, repo.RepositoryName, repo.CloneURL,
						args.DryRun)
					summaryReposCloned = append(summaryReposCloned, cloneResult)
					pbLib.OnSuccess("", fmt.Sprintf("Git clone completed for repository %s, "+
						"with address: %s", repo.RepositoryName, repo.CloneURL))
				}
			} else {
				// Count this repo as skipped, since there was a folder detected with the same name.
				summaryReposSkipped = append(summaryReposSkipped, CloneRepoData{
					Name:     repo.RepositoryName,
					Dir:      specificRepoDir,
					GitURL:   repo.CloneURL,
					IsCloned: false,
				})

				pbLib.OnWarning("", fmt.Sprintf("Git clone skipped, "+
					"repository %s exist with address: %s", repo.RepositoryName, repo.CloneURL))
			}

			pbInstance.Increment()
		}
	}

	// 4.2. Destination directory is not empty, therefore it has content.
	if dirConfig.isWithContent {
		// (Default behavior: neither --skip-if-exists nor --force-if-exists are set)
		// Ignore the existing repositories, and just clone the delta.
		for _, repo := range repositories {
			specificRepoDir := fmt.Sprintf("%s%s", dirConfig.rootDir, repo.RepositoryName)

			// Increment the progress bar.
			pbInstance.UpdateTitle(fmt.Sprintf("Cloning Git Repository ("+
				"existing/with-content dir): %s from address: %s",
				repo.RepositoryName, repo.CloneURL))

			time.Sleep(pbLib.SleepTime)

			if !system.CheckIfDirExists(specificRepoDir) {
				err = system.CreateDir(specificRepoDir)
				if err != nil {
					ux.OutError(fmt.Sprintf("The repository directory %s could not be created. "+
						"Error: %s", specificRepoDir, err.Error()), "", false)

					summaryReposFailed = append(summaryReposFailed, CloneRepoData{
						Name:         repo.RepositoryName,
						Dir:          specificRepoDir,
						GitURL:       repo.CloneURL,
						IsCloned:     false,
						IsSkipped:    false,
						IsError:      true,
						IsDryRun:     false,
						IsOverridden: false,
					})

					pbLib.OnFail("", fmt.Sprintf("Git clone failed for repository %s, "+
						"with address: %s", repo.RepositoryName, repo.CloneURL))
				} else {
					repo.DirCloned = specificRepoDir
					cloneResult := clone(ghSvc, repo.DirCloned, repo.RepositoryName, repo.CloneURL,
						args.DryRun)
					summaryReposCloned = append(summaryReposCloned, cloneResult)
					pbLib.OnSuccess("", fmt.Sprintf("Git clone completed for repository %s, "+
						"with address: %s", repo.RepositoryName, repo.CloneURL))
				}
			} else {
				if args.SkipIfExists {
					// Count this repo as skipped, since there was a folder detected with the same name.
					summaryReposSkipped = append(summaryReposSkipped, CloneRepoData{
						Name:         repo.RepositoryName,
						Dir:          specificRepoDir,
						GitURL:       repo.CloneURL,
						IsCloned:     false,
						IsSkipped:    false,
						IsError:      true,
						IsDryRun:     false,
						IsOverridden: false,
					})

					pbLib.OnWarning("", fmt.Sprintf("Git clone skipped, "+
						"repository %s exist with address: %s", repo.RepositoryName, repo.CloneURL))
				} else {
					if args.ForceIfExist {
						isDeleted, errInDelDir := system.DeleteDir(specificRepoDir, true)
						if errInDelDir != nil {
							ux.OutError(fmt.Sprintf("The repository directory %s could not be deleted. "+
								"Error: %s", specificRepoDir, err.Error()), "", false)

							summaryReposFailed = append(summaryReposFailed, CloneRepoData{
								Name:         repo.RepositoryName,
								Dir:          specificRepoDir,
								GitURL:       repo.CloneURL,
								IsCloned:     false,
								IsSkipped:    false,
								IsError:      true,
								IsDryRun:     false,
								IsOverridden: false,
							})

							pbLib.OnFail("", fmt.Sprintf("Git clone failed for repository %s, "+
								"with address: %s", repo.RepositoryName, repo.CloneURL))
						} else {
							if isDeleted {
								err = system.CreateDir(specificRepoDir)
								if err != nil {
									ux.OutError(fmt.Sprintf("The repository directory %s could not be created. "+
										"Error: %s", specificRepoDir, err.Error()), "", false)

									summaryReposFailed = append(summaryReposFailed, CloneRepoData{
										Name:         repo.RepositoryName,
										Dir:          specificRepoDir,
										GitURL:       repo.CloneURL,
										IsCloned:     false,
										IsSkipped:    false,
										IsError:      true,
										IsDryRun:     false,
										IsOverridden: false,
									})

									pbLib.OnFail("", fmt.Sprintf("Git clone failed ("+
										"force-if-exist is set)"+
										"for repository %s, "+
										"with address: %s", repo.RepositoryName, repo.CloneURL))
								} else {
									repo.DirCloned = specificRepoDir
									cloneResult := clone(ghSvc, repo.DirCloned, repo.RepositoryName, repo.CloneURL,
										args.DryRun)
									summaryReposCloned = append(summaryReposCloned, cloneResult)

									pbLib.OnSuccess("", fmt.Sprintf("Git clone completed for repository %s, "+
										"with address: %s", repo.RepositoryName, repo.CloneURL))

									summaryReposOverridden = append(summaryReposOverridden, CloneRepoData{
										Name:         repo.RepositoryName,
										Dir:          specificRepoDir,
										GitURL:       repo.CloneURL,
										IsCloned:     true,
										IsSkipped:    false,
										IsError:      false,
										IsDryRun:     false,
										IsOverridden: true,
									})
								}
							}
						}

					} else {
						// Count this repo as skipped, since there was a folder detected with the same name.
						summaryReposSkipped = append(summaryReposSkipped, CloneRepoData{
							Name:         repo.RepositoryName,
							Dir:          specificRepoDir,
							GitURL:       repo.CloneURL,
							IsCloned:     false,
							IsSkipped:    true,
							IsError:      false,
							IsDryRun:     false,
							IsOverridden: false,
						})

						pbLib.OnWarning("", fmt.Sprintf("Git clone skipped, "+
							"repository %s exist with address: %s", repo.RepositoryName, repo.CloneURL))
					}
				}
			}

			pbInstance.Increment()
		}
	}

	// Printing summary.
	showSummary(summaryReposCloned, globals.Organization, "Cloned")
	showSummary(summaryReposSkipped, globals.Organization, "Skipped")
	showSummary(summaryReposFailed, globals.Organization, "Failed")
	showSummary(summaryReposOverridden, globals.Organization, "Overridden")
}

func showSummary(data []CloneRepoData, org string, status string) {
	var tableHeaders [][]string
	tableHeaders = append(tableHeaders, []string{"Org", "Repository", "URL", "Status"})

	var tableData [][]string
	for _, item := range data {
		var row [][]string
		row = append(row, []string{org, item.Name, item.GitURL, status})
		tableData = append(tableData, row...)
	}

	tableData = append(tableData, tableData...)

	if status == "Failed" {
		_ = ux.ShowHeader(fmt.Sprintf("Summary of [%s] Repositories", status), "error")
		if len(data) == 0 {
			ux.OutInfo("No repositories failed to clone.", "result")
		}
	}

	if status == "Overridden" {
		_ = ux.ShowHeader(fmt.Sprintf("Summary of [%s] Repositories", status), "warning")
		if len(data) == 0 {
			ux.OutInfo("No repositories were overridden.", "result")
		}
	}

	if status == "Cloned" {
		_ = ux.ShowHeader(fmt.Sprintf("Summary of [%s] Repositories", status), "success")
		if len(data) == 0 {
			ux.OutWarn("No repositories were cloned.", "result")
		}
	}

	if status == "Skipped" {
		_ = ux.ShowHeader(fmt.Sprintf("Summary of [%s] Repositories", status), "warning")
		if len(data) == 0 {
			ux.OutInfo("No repositories were skipped.", "result")
		}
	}

	_ = pterm.DefaultTable.WithHasHeader().WithData(tableData).Render()
}
