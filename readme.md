Gonc – Simple TCP Port Tester (Go)
Description

Gonc est un outil CLI léger écrit en Go permettant de tester la connectivité TCP vers un ou plusieurs ports sur une cible donnée. Il agit comme un mini scanner réseau capable de mesurer le temps de connexion et de vérifier si un port est joignable.

Fonctionnalités
- Test de connexion TCP vers une cible (host:port)
- Support de plusieurs ports en entrée
- Mesure du temps de connexion (latence TCP)

Affichage du statut :
OK si le port est accessible
erreur détaillée sinon

Utilisation simple via arguments CLI
Installation

1. Cloner ou copier le projet
```bash
git clone <repo-url>
cd gonc
```

2. Compiler le projet
```bash
go build
```

Utilisation
Syntaxe
```bash
gonc.exe -host <target> -port <port1,port2,port3> -data <tcp/udp>
```

Exemple
```bash
.\gonc.exe -host google.com -port 80,443,22 -data tcp
Exemple de sortie
OK - google.com:443 took 32ms
OK - google.com:80 took 45ms
Error connecting: dial tcp ... connection refused
```

Principe technique
L’outil repose sur :
- net.Dial() pour établir une connexion TCP/UDP
- time.Now() / time.Since() pour mesurer la latence
- flag pour la gestion des arguments CLI
- strings.Split() pour gérer plusieurs ports

Limitations
- Pas de timeout personnalisé (peut bloquer sur certains ports filtrés)
- Ne vérifie pas la réponse applicative, seulement la connexion TCP

Améliorations possibles
- Ajout de goroutines pour scan parallèle
- Ajout d’un timeout configurable
- Affichage formaté (table, JSON)
- Détection de services (banner grabbing)
