//+build mage

package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/magefile/mage/sh"
	"github.com/mholt/archiver/v3"
	"howett.net/plist"
)

func Build() error {
	err := sh.Rm("Panoptes.saver.zip")

	if err != nil {
		return err
	}

	err = sh.Run("xcodebuild",
		"-project",
		"Panoptes.xcodeproj",
		"-scheme",
		"Panoptes",
		"-configuration",
		"Release",
		"-derivedDataPath",
		"build",
		"clean",
		"build",
	)

	if err != nil {
		return err
	}

	return archiver.Archive([]string{"build/Build/Products/Release/Panoptes.saver"}, "Panoptes.saver.zip")
}

type uploadResponse struct {
	OSVersion      string     `plist:"os-version"`
	SuccessMessage string     `plist:"success-message"`
	Upload         uploadInfo `plist:"notarization-upload"`
}

type uploadInfo struct {
	RequestUUID string `plist:"RequestUUID"`
}

type infoResponse struct {
	OSVersion      string           `plist:"os-version"`
	SuccessMessage string           `plist:"success-message"`
	Info           notarizationInfo `plist:"notarization-info"`
}

type notarizationInfo struct {
	LogFileURL    string `plist:"LogFileURL"`
	Status        string `plist:"Status"`
	StatusMessage string `plist:"Status Message"`
}

func Notarize() error {
	output, err := sh.Output("xcrun",
		"altool",
		"--notarize-app",
		"--primary-bundle-id",
		"com.softwarepunk.Panoptes",
		"--username",
		os.Getenv("NOTARY_USER"),
		"--password",
		"@env:NOTARY_PASSWORD",
		"--file",
		"Panoptes.saver.zip",
		"--output-format",
		"xml",
	)

	if err != nil {
		return err
	}

	var response uploadResponse
	_, err = plist.Unmarshal([]byte(output), &response)

	if err != nil {
		return err
	}

	log.Printf("Upload to notary succeeded, UUID: %s\n", response.Upload.RequestUUID)

	status, err := getNotarizationStatus(response.Upload.RequestUUID)

	if err != nil {
		return err
	}

	for status.Status == "in progress" {
		time.Sleep(10 * time.Second)

		status, err = getNotarizationStatus(response.Upload.RequestUUID)

		if err != nil {
			return err
		}
	}

	log.Printf("Notarization completed with status: '%s'\n", status.Status)

	if status.Status != "success" {
		return fmt.Errorf("NOTARIZATION FAILED WITH: %#v\n\n", status)
	}

	return nil
}

func Staple() error {
	err := sh.Run("xcrun",
		"stapler",
		"staple",
		"build/Build/Products/Release/Panoptes.saver",
	)

	if err != nil {
		return err
	}

	err = sh.Rm("Panoptes.saver.zip")

	if err != nil {
		return err
	}

	return archiver.Archive([]string{"build/Build/Products/Release/Panoptes.saver"}, "Panoptes.saver.zip")
}

func getNotarizationStatus(uuid string) (notarizationInfo, error) {
	output, err := sh.Output("xcrun",
		"altool",
		"--notarization-info",
		uuid,
		"--username",
		os.Getenv("NOTARY_USER"),
		"--password",
		"@env:NOTARY_PASSWORD",
		"--output-format",
		"xml",
	)

	if err != nil {
		return notarizationInfo{}, err
	}

	var response infoResponse
	_, err = plist.Unmarshal([]byte(output), &response)

	if err != nil {
		return notarizationInfo{}, err
	}

	return response.Info, nil
}
