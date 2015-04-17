package main_test

import (
    "strings"
    "strconv"
    "math"
    "os/exec"
    "github.com/onsi/gomega/gexec"

    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
)

var _ = Describe("CLI", func() {
    var pathToCLI string

    BeforeSuite(func() {
        var err error
        pathToCLI, err = gexec.Build("github.com/amitkgupta/cli")
        Ω(err).ShouldNot(HaveOccurred())
    })

    AfterSuite(func() {
        gexec.CleanupBuildArtifacts()
    })

    It("calculates π", func() {
        command := exec.Command(pathToCLI, "pi", "-n", "100000")
        session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
        Ω(err).ShouldNot(HaveOccurred())

        Ω(session.Wait()).Should(gexec.Exit(0))
        resultBytes := session.Wait().Out.Contents()
        resultStr := strings.TrimSpace(string(resultBytes))
        result, err := strconv.ParseFloat(resultStr, 64)
        Ω(err).ShouldNot(HaveOccurred())
        Ω(result).Should(BeNumerically("~", math.Pi, 0.0001))
    })

    It("handles bad input", func() {
        command := exec.Command(pathToCLI, "pi", "-n", "-1")
        session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
        Ω(err).ShouldNot(HaveOccurred())

        Ω(session.Wait()).Should(gexec.Exit(1))
        errBytes := session.Wait().Err.Contents()
        errString := strings.TrimSpace(string(errBytes))
        Ω(errString).Should(ContainSubstring("Invalid n '-1'"))
    })
})
