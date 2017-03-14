// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/google/go-github/github"
	"fmt"
	"github.com/robfig/cron"
	//"golang.org/x/oauth2"
	"context"
	"errors"
)

var (
	currentPRs []*github.PullRequest
	cronJob = cron.New()
	cronExp = "@every 30m"
	githubPR = true
	owner string
	repo string
)

// poll_githubCmd represents the poll_github command
var poll_githubCmd = &cobra.Command{
	Use:   "poll_github",
	Short: "Polls a github repo for PRs or changes",
	Long: `For example:

To poll for PRs: minimalCI poll_github <GITHUB_URL> --cron <CRON_EXP> --githubPR 
To poll a specific branch for new commits: minimalCI poll_github <GITHUB_URL> --cron <CRON_EXP> --branch <BRANCH_NAME>

--githubPR will cause the command to watch for new github Pull Requests.  When a new Pull Request is detected the command specified in minimalCI.yaml will execute.
--cron is required. and is a standard cron expression.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("poll_github called")
		startPoll(cronExp, owner, repo)
		
	},
	
}

func init() {
	RootCmd.AddCommand(poll_githubCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// poll_githubCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// poll_githubCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	
	//cronExp = poll_githubCmd.Flags().String("cron", "", "A cron expression used to determine how often github is polled.")
	//githubPR = poll_githubCmd.Flags().String("githubPR", "", "if set will poll for new github requests")

}


func IsNewPR(PR *github.PullRequest) bool {
    for _, oldPR := range currentPRs {
        if oldPR == PR {
            return false
        }
    }
    return true
}


func GetLatestPRs(owner string, repo string) {
	ctx := context.Background()
	var client = github.NewClient(nil)
	
	// get all pages of results
	
    PRs, resp, err := client.PullRequests.List (ctx, owner, repo, nil)
    if err != nil {
        panic(err)
    }
    if (resp == nil) {
    	err := errors.New("no response")
    	panic(err)
    }

	
	for _, PR := range PRs {
		if IsNewPR(PR) {
			fmt.Printf("new PR detected")
		}
	}
	currentPRs = PRs
}

func startPoll(cronExp string, owner string, repo string) {
	
	cronJob.AddFunc(cronExp, func() { GetLatestPRs(owner, repo) })
	cronJob.Start()
}

func stopPoll() {
	cronJob.Stop()
}	
