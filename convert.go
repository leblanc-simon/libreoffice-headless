package main

import (
    "strings"
    "bytes"
    "log"
    "os/exec"
    "os/user"
    "flag"
    "fmt"
)

func main() {
    filename := flag.String("filename", "", "The filename to convert (without path)")
    path := flag.String("path", "", "The absolute path to the file to convert")

    flag.Parse()

    // Get the current user
    user, err := user.Current()
    if err != nil {
        log.Fatal(err)
    }

    //
    // Check if the path is ending by / (remove it if necessary)
    //
    finalPath := strings.TrimSuffix(*path, "/")

    //
    // Construct the real path of the file to convert
    //
    var filePath bytes.Buffer
    filePath.WriteString("/tmp/")
    filePath.WriteString(*filename)

    //
    // Prepare options
    //
    var environmentFilename bytes.Buffer
    var environmentUid bytes.Buffer
    var environmentGid bytes.Buffer
    var mountDirectory bytes.Buffer

    // - filename
    environmentFilename.WriteString("FILENAME=")
    environmentFilename.WriteString(filePath.String())
    // - uid
    environmentUid.WriteString("UID=")
    environmentUid.WriteString(user.Uid)
    // - gid
    environmentGid.WriteString("GID=")
    environmentGid.WriteString(user.Gid)
    // - mount directory
    mountDirectory.WriteString(finalPath)
    mountDirectory.WriteString(":/tmp")

    // Build command
    cmd := exec.Command(
        "/usr/bin/docker", "run",
        "-e", environmentFilename.String(),
        "-e", environmentUid.String(),
        "-e", environmentGid.String(),
        "-v", mountDirectory.String(),
        "alpine:libreoffice-headless")

    // Execute command
    var out bytes.Buffer
    cmd.Stdout = &out
    err = cmd.Run()

    if err != nil {
        log.Fatalln(err)
    }

    fmt.Println(out.String())
}

