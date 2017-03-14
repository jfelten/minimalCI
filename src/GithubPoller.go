package MinimalCI


import ("github.com/google/go-github/github"
		"fmt"
		"github.com/robfig/cron"
)
var (
	currentPRs []*github.PullRequest
	cron := cron.New()
)

func IsNewPR(PR) bool {
    for _, oldPR := range currentPRs {
        if oldPR == PR {
            return false
        }
    }
    return true
}


func GetLatestPRs(owner, repo) {}
	client := github.NewClient(nil)
	
	opt := &github.PullRequest{
	    ListOptions: github.ListOptions{PerPage: 10},
	}
	// get all pages of results
	var allPRs []*github.PullRequest
	for {
	    PRs, resp, err := client.PullRequests. List (ctx, owner, repo)
	    if err != nil {
	        return err
	    }
	    allPRs = append(allPRs, PRs...)
	    if resp.NextPage == 0 {
	        break
	    }
	    opt.ListOptions.Page = resp.NextPage
	}
	
	for for _, PR := range allPRs {
		if IsNewPR(PR) {
			fmt.Printf("new PR detected")
		}
	}
	currentPRs = allPRs
}

func startPoll(cronExp, owner, repo) {
	
	cron.AddFunc(cronExp, func() { GetLatestPRs(owner, repo) })
	cron.Start()
}

func stopPoll() {
	cron.Stop()
}
}	