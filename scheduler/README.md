# Scheduler

Un gestionnaire de tâches (Todo List) simple et efficace en ligne de commande (CLI), écrit en Go.

## Fonctionnalités

- **Ajout de tâches** : Enregistrez rapidement des tâches avec une description.
- **Liste des tâches** : Affichez des tâches en attente avec un calcul du temps écoulé (ex: "2 hours ago").
- **Marquage comme terminé** : Validez des tâches une fois accomplies.
- **Suppression** : Supprimez les tâches inutiles.
- **Persistance des données** : Les tâches sont sauvegardées localement dans un fichier CSV (`storage/files/data.csv`).
- **Verrouillage de fichier** : Utilise `flock` pour garantir l'intégrité des données lors des écritures.

## Installation

Assurez-vous d'avoir [Go](https://go.dev/doc/install) installé sur votre machine.

1. Clonez le dépôt :
   ```bash
   git clone <url-du-depot>
   cd scheduler
   ```

2. Installez les dépendances :
   ```bash
   go mod download
   ```

3. Compilez l'application :
   ```bash
   go build -o scheduler main.go
   ```

## Utilisation

L'application s'utilise via différentes commandes :

### Ajouter une tâche
```bash
./scheduler add "Acheter du pain"
```

### Lister les tâches en attente
```bash
./scheduler list
```

### Lister toutes les tâches (y compris terminées)
```bash
./scheduler list -a
```

### Marquer une tâche comme terminée
Utilisez l'ID de la tâche affiché lors du `list`.
```bash
./scheduler complete 1
```

### Supprimer une tâche
```bash
./scheduler delete 1
```

## Structure du projet

- `main.go` : Point d'entrée de l'application.
- `cli/` : Gestion des arguments et de l'interface utilisateur.
- `todo/` : Logique métier des tâches.
- `storage/` : Gestion de la persistance en CSV et verrouillage des fichiers.

## Dépendances

- [timediff](https://github.com/mergestat/timediff) : Pour l'affichage formaté du temps écoulé.
