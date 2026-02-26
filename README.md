# Projet Docker

Ce projet consiste à conteneuriser une application full-stack simple afin d’appliquer les notions Docker vues en cours. L’application permet d’ajouter et de lister des utilisateurs avec leur age stockés dans une base de données PostgreSQL.

---

## Sommaire

1. Prérequis
2. Notions Docker utilisées
3. Architecture technique
4. Quick Start
5. Tester l'application
6. Persistance des données
7. Contributeurs

---

## 1. Prérequis

Avant de lancer le projet, assurez-vous d’avoir installé :

- Docker
- Docker Compose (généralement inclus avec Docker Desktop)

Vérifier l'installation :

```bash
docker --version
docker compose version
```

## 2. Notions Docker utilisées

Les notions Docker appliquées dans ce projet :

### Dockerfile

Chaque service (front et back) possède un Dockerfile permettant de construire son image.

### Multi-stage build

Utilisé pour :

- réduire la taille des images
- séparer les étapes de build et d'exécution

### Docker Compose

Le projet utilise docker compose pour gérer plusieurs conteneurs :

- gestion de plusieurs services
- création d’un réseau interne
- gestion d'une volume pour la persistance
- gestion des variables sensibles via `.env` et `.secret`

### Bonnes pratiques Docker

1. Users non-root
   Les conteneurs utilisent un utilisateur non root pour des raisons de sécurité.

2. **.dockerignore**

Permet de :

- exclure les fichiers sensibles
- réduire la taille du contexte de build
- accélérer la construction des images

3. **alpine:3.20**

Bonne pratique d'un point de vue sécurité

## 3. Architecture technique

L'application est composée de 3 services Docker.

### Frontend

- Application web simple (HTML / CSS / JS)
- Permet :
  - d'ajouter un utilisateur
  - de lister les utilisateurs existants

### Backend

API Go qui :

- reçoit les requêtes HTTP du frontend
- valide les données
- interagit avec la base de données

### Base de données

Base de données relationnelle PostGreSQL liée à une volume Docker pour la persistance de données

## 4. Quick Start

Cloner le repo du projet:

```bash
git clone https://github.com/DoMinhCat/Docker-Project.git
cd Docker-Project
```

Construire les images, les containeurs et lancer l'application:

```bash
docker compose up
```

### 5. Tester l'application

- Aller sur: <http://localhost:3002/>
- Cliquer sur le bouton: `List All Users`
- Vous verrez 2 utilisateurs déjà présent (créés par fichier `.sql`):

```Code
Roman Lenoir 18
titi 51
```

- Remplir le formulaire avec un nom et un age puis cliquer sur `Submit`
- Arrêter les containeurs (ctrc + c ou via Docker Desktop)
- Relancer les containeurs
- Vous verrez que l'utilisateur que vous avez ajouté est toujour présent, ce qui confirme que les données sont persistées grâce au volume Docker PostgreSQL.

## 3. Contributeurs

- Arnaud
- Ayoub
- Minh Cat DO

ESGI 2025 - 2026
