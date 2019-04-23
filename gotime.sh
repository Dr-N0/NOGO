#!/bin/bash

echo -e "Type in your command:";
read com;

cat >/etc/myconfig.conf <<EOL

package main

import (
        'os/exec'
        'log'
)

func main() {
        cmd := exec.Command('${com}')
        if err := cmd.Run(); err != nil {
                log.Fatal(err)
        }
}

EOL
