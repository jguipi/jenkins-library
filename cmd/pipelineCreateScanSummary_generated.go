// Code generated by piper's step-generator. DO NOT EDIT.

package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/SAP/jenkins-library/pkg/config"
	"github.com/SAP/jenkins-library/pkg/log"
	"github.com/SAP/jenkins-library/pkg/telemetry"
	"github.com/spf13/cobra"
)

type pipelineCreateScanSummaryOptions struct {
	OutputFilePath string `json:"outputFilePath,omitempty"`
}

// PipelineCreateScanSummaryCommand Collect scan result information anc create a summary report
func PipelineCreateScanSummaryCommand() *cobra.Command {
	const STEP_NAME = "pipelineCreateScanSummary"

	metadata := pipelineCreateScanSummaryMetadata()
	var stepConfig pipelineCreateScanSummaryOptions
	var startTime time.Time

	var createPipelineCreateScanSummaryCmd = &cobra.Command{
		Use:   STEP_NAME,
		Short: "Collect scan result information anc create a summary report",
		Long: `This step allows you to create a summary report of your scan results.

It is for example used to create a markdown file which can be used to create a GitHub issue.`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			startTime = time.Now()
			log.SetStepName(STEP_NAME)
			log.SetVerbose(GeneralConfig.Verbose)

			path, _ := os.Getwd()
			fatalHook := &log.FatalHook{CorrelationID: GeneralConfig.CorrelationID, Path: path}
			log.RegisterHook(fatalHook)

			err := PrepareConfig(cmd, &metadata, STEP_NAME, &stepConfig, config.OpenPiperFile)
			if err != nil {
				log.SetErrorCategory(log.ErrorConfiguration)
				return err
			}

			if len(GeneralConfig.HookConfig.SentryConfig.Dsn) > 0 {
				sentryHook := log.NewSentryHook(GeneralConfig.HookConfig.SentryConfig.Dsn, GeneralConfig.CorrelationID)
				log.RegisterHook(&sentryHook)
			}

			return nil
		},
		Run: func(_ *cobra.Command, _ []string) {
			telemetryData := telemetry.CustomData{}
			telemetryData.ErrorCode = "1"
			handler := func() {
				config.RemoveVaultSecretFiles()
				telemetryData.Duration = fmt.Sprintf("%v", time.Since(startTime).Milliseconds())
				telemetryData.ErrorCategory = log.GetErrorCategory().String()
				telemetry.Send(&telemetryData)
			}
			log.DeferExitHandler(handler)
			defer handler()
			telemetry.Initialize(GeneralConfig.NoTelemetry, STEP_NAME)
			pipelineCreateScanSummary(stepConfig, &telemetryData)
			telemetryData.ErrorCode = "0"
			log.Entry().Info("SUCCESS")
		},
	}

	addPipelineCreateScanSummaryFlags(createPipelineCreateScanSummaryCmd, &stepConfig)
	return createPipelineCreateScanSummaryCmd
}

func addPipelineCreateScanSummaryFlags(cmd *cobra.Command, stepConfig *pipelineCreateScanSummaryOptions) {
	cmd.Flags().StringVar(&stepConfig.OutputFilePath, "outputFilePath", `scanSummary.md`, "Defines the filepath to the target file which will be created by the step.")

}

// retrieve step metadata
func pipelineCreateScanSummaryMetadata() config.StepData {
	var theMetaData = config.StepData{
		Metadata: config.StepMetadata{
			Name:        "pipelineCreateScanSummary",
			Aliases:     []config.Alias{},
			Description: "Collect scan result information anc create a summary report",
		},
		Spec: config.StepSpec{
			Inputs: config.StepInputs{
				Parameters: []config.StepParameters{
					{
						Name:        "outputFilePath",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
				},
			},
		},
	}
	return theMetaData
}
