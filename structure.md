# Standard Go Project Layout

### `/internal`

Applications privées et bibliothèques de code. C'est le code que vous ne souhaitez pas voir importé dans d'autres applications ou bibliothèques. Notez que ce pattern est imposé par le compilateur Go lui-même (voir les [`release notes`](https://golang.org/doc/go1.4#internalpackages) de Go 1.4 pour plus de détails). Vous n'êtes pas limité à un seul dossier `internal` de haut niveau, mais vous pouvez en avoir plusieurs à n'importe quel niveau de l'arborescence de votre projet.

Vous pouvez également ajouter un peu de structure dans vos packages internes pour séparer le code partagé et non partagé. Ce n'est pas du tout obligatoire (surtout pour les petits projets), mais il est intéressant d'avoir des indices visuels indiquant l'utilisation prévue d'un package. Le code de votre application peut aller dans un dossier `/internal/app` (p. ex., `/internal/app/myapp`) tandis que le code partagé par les applications peut se retrouver dans un dossier `/internal/pkg` (p. ex., `/internal/pkg/myprivlib`).

### `/pkg`

Placez-y le code qui peut être réutilisé par les applications externes (p. ex., `/pkg/mypubliclib`). D'autres projets peuvent importer ces bibliothèques et s'attendent donc à ce qu'elles soient fonctionnelles, pensez y donc à deux fois avant de mettre du code dans ce dossier :-) Utiliser le dossier `internal` est une manière plus adéquate de garder vos packages privés et non importables car c'est intégré au compilateur Go. Le dossier `/pkg` est néanmoins une bonne manière d'indiquer que le code contenu dans ce dossier peut être utilisé par les autres utilisateurs sans problème. L'article de blog de Travis Jeffery [`I'll take pkg over internal`](https://travisjeffery.com/b/2019/11/i-ll-take-pkg-over-internal/) explique plus en détail les différences entre les dossier `pkg` et `internal` et quand il fait sens de les utiliser.

C'est également une manière de regrouper tout votre code Go au même endroit lorsque votre dossier racine comporte de nombreux composants et dossiers non-Go, permettant plus facilement de lancer les différents outils Go, tel que mentionné dans les conférences suivantes : [`Best Practices for Industrial Programming`](https://www.youtube.com/watch?v=PTE4VJIdHPg) lors de GopherCon EU 2018, [GopherCon 2018: Kat Zien - How Do You Structure Your Go Apps](https://www.youtube.com/watch?v=oL6JBUk6tj0) et [GoLab 2018 - Massimiliano Pippi - Project layout patterns in Go](https://www.youtube.com/watch?v=3gQa1LWwuzk)).

Voir le dossier [`/pkg`](pkg/README.md) pour découvrir quels projets Go populaires utilisent cette architecture de projet. C'est un pattern plutôt commun, mais qui n'est pas accepté de manière universelle, et certaines personnes de la communauté Go ne le recommandent pas.

Vous n'êtes pas obligés de l'utiliser si votre projet est petit et si l'ajout d'un niveau de plus n'ajoute pas vraiment de valeur (sauf si vous y tenez vraiment :-)). Il est temps d'y penser lorsque votre projet commence à prendre de l'ampleur et que votre dossier racine est encombré (surtout si vous avez beaucoup de composants non-Go)

### `/vendor`

Les dépendances de votre application (gérées manuellement ou via votre gestionnaire de dépendances favori tel que la fonctionnalité incluse dans les [`Go Modules`](https://github.com/golang/go/wiki/Modules)). La commande `go mod vendor` créera un dossier `/vendor` pour vous. Notez que vous devrez peut-être utiliser le flag `-mod=vendor` avec votre commande `go build` si vous n'utilisez pas Go 1.14 qui le définit par défaut.

Ne commitez pas vos dépendances si vous développez une bibliothèque.

Depuis sa version [`1.13`](https://golang.org/doc/go1.13#modules), Go active la fonctionnalité de proxy de module (en utilisant [`https://proxy.golang.org`](https://proxy.golang.org) comme serveur de proxy par défaut). Plus d'infos [`ici`](https://blog.golang.org/module-mirror-launch) afin de définir si cela correspond à votre obligations et contraintes. Si c'est le cas, vous n'aurez pas besoin du dossier `vendor`.

## Les répertoires d'application de services

### `/api`

Spécifications OpenAPI/Swagger, fichiers de schémas JSON, fichiers de définitions de protocoles.

Voir le dossier [`/api`](api/README.md) pour des examples.

## Les répertoires d'application web

### `/web`

Les composants spécifiques aux applications web : assets statiques, templates serveurs et SPAs.

## Les répertoire communs aux applications

### `/configs`

Templates de fichiers de configuration ou configurations par défaut.

Ajoutez vos templates `confd` ou `consul-template` dans ce répertoire.

### `/init`

Initialisation du système (systemd, upstart, sysv) et configurations des administrateurs/superviseurs de process (runit, supervisord).

### `/scripts`

Scripts permettant différentes opérations telles que le build, l'installation, des analyses, ...

Ces scripts permettent de garder le Makefile du dossier racine réduit et simple (p. ex., [`https://github.com/hashicorp/terraform/blob/main/Makefile`](https://github.com/hashicorp/terraform/blob/main/Makefile)).

Voir le dossier [`/scripts`](scripts/README.md) pour des exemples.

### `/build`

Packaging et Intégration Continue.

Ajoutez vos scripts et configurations de cloud (AMI), conteneur (Docker), OS (deb, rpm, pkg) et package dans le dossier `/build/package`.

Placez vos scripts et configurations de CI (travis, circle, drone) dans le dossier `/build/ci`. Prenez garde au fait que certains outils de CI (p. ex., Travis CI) sont très contraignants vis à vis de l'emplacement de leurs fichiers de configuration. Essayez donc, lorsque c'est possible, de créer des liens entre le dossier `/build/ci` et les endroits où les outils de CI s'attendent à trouver ces fichiers.

### `/deployments`

Templates et configurations pour les IaaS, PaaS, système et l'orchestration de conteneurs (docker-compose, kubernetes/helm, mesos, terraform, bosh). Sur certains projets (principalement les applications déployées via Kubernetes) ce dossier s'appelle `/deploy`.

### `/test`

Applications et données de tests externes additionnels. Vous pouvez structurer le dossier `/test` de la façon qui vous convient le mieux. Pour des projets plus importants, il fait sens d'utiliser un sous-dossier `data`. Vous pouvez par exemple utiliser `/test/data` ou `/test/testdata` si vous souhaitez que Go ignore ce dossier. Go ignore également les dossiers ou fichiers commençant par "." ou "\_", ce qui vous donne plus de flexibilité en terme de nommage de votre dossier de données de test.

Voir le dossier [`/test`](test/README.md) pour des exemples

## Autres répertoires

### `/docs`

Documents utilisateurs et design (en plus de votre documentation générée GoDoc)

Voir le dossier [`/docs`](docs/README.md) pour des exemples

### `/tools`

Outils de support du projet. Ces scripts peuvent importer du code des dossier `/pkg` et `/internal`

Voir le dossier [`/tools`](tools/README.md) pour des exemples

### `/examples`

Exemples de vos applications et/ou de vos bibliothèques publiques

Voir le dossier [`/examples`](examples/README.md) pour des exemples

### `/third_party`

Outils d'aide externe, code forké et autres utilitaires tierces (p. ex., Swagger UI).

### `/githooks`

Hooks Git.

### `/assets`

D'autres assets qui sont utilisés dans votre dépôt (images, logos, etc).

### `/website`

C'est là que vous placez les données du site de votre projet si vous n'utilisez pas GitHub pages.

Voir le dossier [`/website`](website/README.md) pour des exemples

## Les répertoires que vous ne devriez pas avoir

### `/src`

Certains projets Go comportent un dossier `src` mais cela arrive en général lorsque les développeurs viennent du monde de Java où c'est une pratique habituelle. Faites tout votre possible pour ne pas adopter ce pattern Java. Vous n'avez vraiment pas envie que votre code Go ou vos projets Go ressemblent à du Java :-)

Ne confondez pas le répertoire `/src` à la racine avec le répertoire `/src` utilisé par Go pour gérer ses espaces de travail comme décrit dans [`How to Write Go Code`](https://golang.org/doc/code.html). La variable d'environnement `$GOPATH` pointe vers votre espace de travail courant (par défault il pointe vers `$HOME/go` sur les systèmes non-Windows). Cet espace de travail inclut les dossiers `/pkg`, `/bin` et `/src`. Votre projet en lui-même va se retrouver dans un sous-dossier de `/src`, donc si vous avez un dossier `/src` dans votre projet le chemin vers celui-ci ressemblera à ceci : `/some/path/to/workspace/src/your_project/src/your_code.go`. Notez qu'à partir de Go 1.11 il est possible d'avoir votre projet en dehors de votre `GOPATH` mais cela ne veut toujours pas dire que c'est une bonne idée d'utiliser le dossier `/src`

## Badges

-   [Go Report Card](https://goreportcard.com/) - Scanne votre code avec les commandes `gofmt`, `go vet`, `gocyclo`, `golint`, `ineffassign`, `license` and `misspell`. Remplacez `github.com/golang-standards/project-layout` avec l'url de votre projet.

-   ~~[GoDoc](http://godoc.org) - Fournit une version en ligne de votre documentation générée GoDoc. Modifiez le lien pour qu'il pointe vers votre projet.~~

-   [Pkg.go.dev](https://pkg.go.dev) - Pkg.go.dev est la nouvelle destination pour la découverte de Go et sa documentation. Vous pouvez créer une badge en utilisant [l'outil de création de badge](https://pkg.go.dev/badge).

-   Release - Il indique la dernière version de votre projet. Modifiez le lien GitHub pour qu'il pointe vers votre projet.

[![Go Report Card](https://goreportcard.com/badge/github.com/golang-standards/project-layout?style=flat-square)](https://goreportcard.com/report/github.com/golang-standards/project-layout)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/golang-standards/project-layout)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/golang-standards/project-layout)](https://pkg.go.dev/github.com/golang-standards/project-layout)
[![Release](https://img.shields.io/github/release/golang-standards/project-layout.svg?style=flat-square)](https://github.com/golang-standards/project-layout/releases/latest)

## Notes

Un template de projet moins générique avec du code, des script et des configs réutilisables est en cours de réalisation.
