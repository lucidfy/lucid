# LUCID

Inspired by Laravel / Symfony structure but written in Go!
This project is still *under development*, there's no released tag yet!

## Installation via Docker

Just execute `./setup-docker` and it should build a container **lucid-container**

The docker image consumes these ports 8080 for lucid and 8081 for svelte-kit, however these ports were internal, it forward back to your docker host under these ports 8080 -> **8330**, 8081 -> **8331**. You can verify this by running `docker ps -a`

Therefore, try to open your browser and access http://localhost:8330 for lucid and http://localhost:8331 for svelte-kit

## Security Concerns

If you found any security concerns, please send a direct email to **daison12006013@gmail.com** the title of the email should have at least a word "Lucid". Thank you!
