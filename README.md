# **go-web-scraper**

Ce projet est un script de **web scraping** en **Go** qui extrait des informations de produits à partir d'un site e-commerce, tout en respectant les règles définies dans le fichier `robots.txt` du site cible.

---

## **Points clés**

### **1. Respect de robots.txt**
Le script lit et prend en compte les restrictions d'accès et de fréquence indiquées dans le fichier `robots.txt` (par exemple, les chemins interdits ou le délai entre les requêtes).

### **2. Extraction via sitemap.xml**
Il utilise le fichier `sitemap.xml` pour récupérer efficacement la liste des **URLs produits** à traiter.

### **3. Simplicité et éthique**
Le script évite les requêtes massives en parallèle et applique un **délai entre chaque requête** pour ne pas surcharger le serveur.

---
<img width="802" height="447" alt="image" src="https://github.com/user-attachments/assets/698e262d-0e7d-4be3-9def-841dde8ceec5" />
