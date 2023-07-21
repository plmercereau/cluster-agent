# Cluster agent

Agent to automatically connect a new node to a nearby Pacemaker cluster.

## TODO

-   Build & release
-   Install:
    -   run the agent as a cluster systemd service
    -   make sure the agent runs as an user with the right permissions:
        -   run `pcs`
        -   change the password of the `hacluster` user (`sudoers` line)
-   Wrap this up into a NixOS package
