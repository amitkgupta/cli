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

        resultBytes := session.Wait().Out.Contents()
        resultStr := strings.TrimSpace(string(resultBytes))
        result, err := strconv.ParseFloat(resultStr, 64)
        Ω(err).ShouldNot(HaveOccurred())
        Ω(result).Should(BeNumerically("~", math.Pi, 0.0001))
    })
})
