# Projet SI - Reverse shell et phishing
## Contexte
Nous avons décider de lier l'infra et la sécu pour ce projet et de créer un programme en go de type command and control qui nous permeterai de gérer plusieurs machines d'un réseau simultanément
    
## Prérequis : 
* Une machine windows sur le réseau attaqué
* une machine windows 64 bits
* une machine linux avec python d'installé
* Une machine distante du réseau avec go d'installé
* Toutes ces machines sont atteignables sur le réseau, elles ont une ip propre.
* Tous les antivirus et pare-feux du réseau sont désactivés pour la démo
## Démarche
* **1** : Modifier l'adresse IP dans le fichier de contrôle et shell par l'adresse IP de la machine windows distante qui controlera les machines infectées
* **2** : Compiler le shell.go en shell.exe dans l'architecture des pc visés ( Ex: 64bits ) avec `go build .\shell.go`
* **3** : Lancer le serveur command control avec le port par défaut 8080 `go run .\Control.go -l -p 8080`
* **4** : Lancer la machine linux, dans un dossier avec le shell.exe, saisir la commande : `python -m SimpleHTTPServer 8080`
* **5** : Modifier le programme vba de la macro avec l'adresse ip publique de la machine linux
* **6** : Envoyer un mail aux victimes avec le fichier excel en `.xlsm`
* **7** : Attendre que les victimes ouvrent le excel avec les macros
* **8** : prendre le controle de la machine grâce au reverse shell
## Explication : 
La victime reçoit un email du gouvernement lui proposant d'analyser son revenu pour obtenir une réduction d'impots, elle télécharge et ouvre le fichier excel et active les macros a cause d'un manque de sécurité. 
A l'activation des macros, un programme vba se lance, télécharge et lance un implant dans le pc de la victime à l'adresse `C:\Program Files\Internet Explorer`.
Le fichier est hébergé sur la machine linux,
Nous avons donc un accès a la machine grâce à ce reverse shell.
## Axes d'amélioration
* Un bypass des antivirus
* Gain de privilège sur la machine
* Gestion multi-connexion
