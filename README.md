# LibreOffice Headless via Alpine

## Installation

Installation sur Ubuntu 16.04

```bash
#
# Installation de docker
#
sudo apt-get update
sudo apt-get install \
    apt-transport-https \
    ca-certificates \
    curl \
    software-properties-common
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
sudo apt-key fingerprint 0EBFCD88
sudo add-apt-repository \
   "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
   $(lsb_release -cs) \
   stable"
sudo apt-get update
sudo apt-get install docker-ce

#
# Ajout du group à l'utilisateur
# !!! Remplacer $USER par l'utilisateur voulu si ce n'est pas l'utilisateur courant
#
sudo usermod -aG docker $USER

# Il faut se relogger entièrement suite à l'ajout

#
# Récupération du dépôt 
#
git clone https://github.com/leblanc-simon/libreoffice-headless.git
cd libreoffice-headless

#
# Construction de l'image
#
docker build -t alpine:libreoffice-headless .
```

## Utilisation

```bash
docker run -e FILENAME=/tmp/document.docx -e UID=1000 -e GID=1000 -v /local-path/tmp:/tmp alpine:libreoffice-headless
```

## Configuration

* `FILENAME` : le chemin vers le fichier à convertir. Il s'agit du chemin pour le container Docker. Il faut donc faire la liaison via l'option `-v` entre le chemin de l'hôte et le chemin vers le container
* `UID` : L'UID à utiliser pour la génération du fichier PDF (logiquement l'UID de l'utilisateur courant)
* `GID` : Le GID à utiliser pour la génération du fichier PDF (logiquement le GID de l'utilisateur courant)

## Auteurs

* [SFoxDev](https://github.com/sfoxdev/docker-unoconv-alpine) (auteur original) 
* Simon Leblanc <contact@leblanc-simon.eu> (reprise du script de SFoxDev)

