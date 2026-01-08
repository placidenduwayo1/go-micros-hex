# HexaShop (Microservices + Architecture Hexagonale)

Projet exemple de microservices structurés en **architecture hexagonale (Ports & Adapters)**.  
Objectif : isoler le **domaine** (métier) des détails techniques (HTTP, DB, messaging), pour faciliter les tests, l’évolution et le remplacement des adapters.

---

## 🧭 Vision

- **Microservices** : chaque service est autonome (code, DB, déploiement).
- **Hexagonal** : le domaine ne dépend de rien.
- **DDD-friendly** : bounded contexts clairs (Customer, Order, Product, Payment…).

---

## 🧱 Microservices

- `customer-service` : gestion des clients et leurs adresses
- `product-service` : catalogue de produits et leur quantité en stock
- `order-service` : commandes, lignes de commande effectuées par les customers, 
- `payment-service` : service de paiements
- `gateway-proxy` : API Gateway point d’entrée dans l'univers des microservices de l'app

---

## 🗂️ Structure `customer-microservice`  en archi hexagonale

customer-microservice/
├── cmd/
│   └── api/
│       ├── main.go                    # composition root (wiring)
│       ├── routes.go                  # register routes (gin/nethttp)
│       └── container.go               # build dependencies (db, repos, usecases, handlers)
│
├── internal/
│   ├── domain/                        # ✅ OBJETS MÉTIER (purs)
│   │   ├── customer.go                # objet métier Customer
│   │   ├── address.go                 # objet métier Address 
│   │   ├── validators/
│   │   │   ├── email.go               # ex: validation email
│   │   │   └── zip_code.go             # ex: validation ZipCode
│   │   │   └── phone_num.go             # ex: validation phoneNum
│   │   └── errors.go                  # erreurs métier (ErrInvalid..., etc.)
│   │
│   ├── application/                   # ✅ USE CASES + PORTS
│   │   ├── ports/
│   │   │   ├── in/                    # le microservice expose les ports d'entrée
│   │   │   │   ├── customer_uc.go     # Create/Get/Update/Delete Customer
│   │   │   │   └── address_uc.go      # Create/Get/Update/Delete Address
│   │   │   └── out/                   # ce dont l’app a besoin pour envoyer à l'extérieur
│   │   │       ├── customer_service.go   # interface CustomerService
│   │   │       └── address_service.go    # interface AddressService
│   │   │
│   │   └── usecase/                   # impl des ports d'entrée (in)
│   │       ├── customer_service.go    # CustomerServiceimpln (usecase customer)
│   │       └── address_service.go     # AddressServiceimpl (usecase address)
│   │
│   ├── infrastructure/                # ✅ ADAPTERS (l'exterieur)
│   │   ├── web/
│   │   │   └── http/
│   │   │       ├── handlers/
│   │   │       │   ├── customer_handler.go
│   │   │       │   └── address_handler.go
│   │   │       ├── dtos/              # ✅ DTOs API
│   │   │       │   ├── customer_request.go
│   │   │       │   ├── customer_response.go
│   │   │       │   ├── address_request.go
│   │   │       │   └── address_response.go
│   │   │       └── mappers/           # ✅ DTO ⇄ DOMAIN
│   │   │           ├── customer_mapper.go
│   │   │           └── address_mapper.go
│   │   │
│   │   ├── persistence/
│   │   │   └── postgres/
│   │   │       ├── db.go              # connection, ping
│   │   │       ├── models/            # ✅ OBJETS BDD (Row models)
│   │   │       │   ├── customer_row.go
│   │   │       │   └── address_row.go
│   │   │       ├── mappers/           # ✅ DOMAIN ⇄ DB
│   │   │       │   ├── customer_mapper.go
│   │   │       │   └── address_mapper.go
│   │   │       └── repositories/      # impl des output ports
│   │   │           ├── customer_repo.go
│   │   │           └── address_repo.go
│   │   │
│   │   └── clock/                     # ex: time provider (optionnel)
│   │       └── system_clock.go
│   │
│   └── config/
│       ├── config.go                  # env vars -> Config struct
│       └── logger.go                  # zap/logrus/std log
│
└── migrations/
    ├── 001_create_addresses.sql
    └── 002_create_customers.sql
