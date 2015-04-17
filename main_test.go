package main_test

import (
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
        command := exec.Command(pathToCLI, "pi", "-n", "1000:0")
        session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
        Ω(err).ShouldNot(HaveOccurred())

        resultStr := session.Wait().Buffer().Contents()
        result, err := strconv.Atoi(string(resultStr))
        Ω(err).ShouldNot(HaveOccurred())
        Ω(result).Should(BeNumerically("~", math.Pi, 0.0001))
    })
})
