// Main package for the mongotop tool.
package main

import (
	"github.com/mongodb/mongo-tools/common/db"
	"github.com/mongodb/mongo-tools/common/log"
	commonopts "github.com/mongodb/mongo-tools/common/options"
	"github.com/mongodb/mongo-tools/mongotop"
	"github.com/mongodb/mongo-tools/mongotop/options"
	"github.com/mongodb/mongo-tools/mongotop/output"
	"os"
	"strconv"
	"time"
)

const (
	// the default sleep time, in seconds
	DEFAULT_SLEEP_TIME = 1
)

func main() {

	// initialize command-line opts
	opts := commonopts.New("mongotop", "<options> <sleeptime>")

	// add mongotop-specific options
	outputOpts := &options.Output{}
	opts.AddOptions(outputOpts)

	extra, err := opts.Parse()
	if err != nil {
		log.Logf(log.Always, "error parsing command line options: %v", err)
		os.Exit(1)
	}

	// print help, if specified
	if opts.PrintHelp(false) {
		return
	}

	// print version, if specified
	if opts.PrintVersion() {
		return
	}

	// pull out the sleeptime
	// TODO: validate args length
	sleeptime := DEFAULT_SLEEP_TIME
	if len(extra) > 0 {
		sleeptime, err = strconv.Atoi(extra[0])
		if err != nil {
			log.Logf(log.Always, "bad sleep time: %v", extra[0])
			os.Exit(1)
		}
	}

	// create a session provider to connect to the db
	sessionProvider, err := db.InitSessionProvider(*opts)
	if err != nil {
		log.Logf(log.Always, "error initializing database session: %v", err)
		os.Exit(1)
	}

	// instantiate a mongotop instance
	top := &mongotop.MongoTop{
		Options:         opts,
		OutputOptions:   outputOpts,
		Outputter:       &output.TerminalOutputter{},
		SessionProvider: sessionProvider,
		Sleeptime:       time.Duration(sleeptime) * time.Second,
		Once:            outputOpts.Once,
	}

	// kick it off
	if err := top.Run(); err != nil {
		log.Logf(log.Always, "error running mongotop: %v", err)
		os.Exit(1)
	}
}
