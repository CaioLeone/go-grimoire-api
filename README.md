# API REST - Go Grimório API
## Conheça o Projeto

O ***Go Grimoire API*** é um projeto prático em Go com foco na construção de uma API RESTful com temática de fantasia, focado no gerenciamento de ***Magias*** e ***Criaturas***. 

O projet foi cirado com o objetivo de evoluir de um CRUD simples para uma arquitetura mais organizada e escalável, aplicando conceitos modernos de desenvolvimento backend utilziando Go.

## Objetivos de Aprendizado
* Desenvolvimento de APIs RESTful com Go
* Trabalhar com ```net/http```
* Manipulação de JSON
* Uso de router ```chi```
* Estruturação de projetos em Go
* Clean Architecture
* Uso de interfaces
* Middlewares
* Persistência com SQLite 
* Queries dinâmicas (filtros)
* Paginação e ordenação
* Separação de responsabilidade
* Relacionamentos enre entidades

## Tecnologias Utilizadas
* Go
* Chi Router
* SQLite
* Swagger
* UUID
* JSON
* REST API
* database/sql

## Estrutura do Projeto
```
├── .dist
├── api
│   ├── domain
│   ├── handler
│   ├── repository
│   └── usecase
└── docs
```

## Arquitetura

O projeto segue uma estrutura inspirada em Clean Architecture
```
/api
  /domain     -> entidades
  /usecase    -> regras de negocio
  /repository -> acesso a dados (memory / sqlite)
  /handler    -> camada HTTP
```

Fluxo:

```Handler -> Usecase -> Repository -> Banco```

## Domínio da Aplicação

A API será responsável por gerenciar:

* Magias
* Criaturas

## Estrutura das Entidades
1. ***Magia (Spell)***
```
{
  "id": "", // UUID, obrigatório
  "name": "Bola de Fogo",
  "description": "Uma explosão flamejante devastadora",
  "mana_cost": 50,
  "element": "Fogo"
  "type": "Ataque"
  "power": "15"
}
```

* Regras:

  1. name: obrigatório (mín. 2)
  2. description: obrigatório (mín. 2)
  3. mana_cost: obrigatório (> 0)
  4. element: obrigatório (ex: Fogo, Gelo, Arcano)
  5. type: obrigatorio (ex: Ataque, defesa, cura, maldição)
  6. power: obrigatorio (> 0)

1. ***Criatura (Creature)***
```
{
  "id": "", // UUID, obrigatório
  "name": "Dragão Vermelho",
  "description": "Uma criatura lendária que cospe fogo",
  "hp": 500,
  "attack": 80,
  "defense": 60,
  "spells": []
}
```
* Regras:
  1. name: obrigatório (mín. 2)
  2. description: obrigatório (mín. 2)
  3. hp, attack, defense: obrigatórios (> 0)

* Relacionamento:
```
{
  "id": "creature_id",
  "spells":[
    {
      "id": "spell_id",
      "name": "bola de fogo"
    }
  ]
}
```

## Persistencia

1. Em memória
  * Uso de ```map``` para simular banco

2. SQLite
  * ```database/sql```
  * driver: ```github.com/mattn/go-sqlite3```

3. Tabelas:
  * ```spells```
  * ```creatures```
  * ```creature_spells``` (Relacionamento)

## Endpoints da API
### Magias
```
POST    /api/spell
GET     /api/spell
GET     /api/spell/{id}
PUT     /api/spell/{id}
DELETE  /api/spell/{id}
```
* Filtros de magia
```
GET /api/spell?element=fire
GET /api/spell?name=fire
GET /api/spell?sort=mana_cost&order=desc
GET /api/spell?page=1&limit=10
```
### Criaturas
```
POST    /api/creature
GET     /api/creature
GET     /api/creature/{id}
PUT     /api/creature/{id}
DELETE  /api/creature/{id}
```
* Ensinar Magia
```
POST /api/creature/{id}/teach/{spellId}
```
* Filtros de Creaturas
```
GET /api/creature?name=dragon
GET /api/creature?attack=50
GET /api/creature?sort=attack&order=desc
GET /api/creature?page=1&limit=10
```

## Paginação
```GET /api/spell?page=1%limit=10```

## Ordenação
```
?sort=name&order=asc
?sort=mana_cost&order=desc
```

## Middlewares
* Logging
* Recovery
* RequestID

## Padrão de Erro

Todas as respostas de erro devem seguir:
```
{
  "error": "mensagem descritiva"
}
```

## Testes Automatizados
O projeto possui testes automatizados para a camada de Usecase:

### Cobertura Atual
* Create Spell
* Update Spell
* Delete Spell
* Get Spell By ID
* Get Spell By Element
* Get All Spells
* Create Creature
* Update Creature
* Delete Creature
* Get Creature By ID
* Teach Spell
* Validações de regras de negócio

### Executar Testes
```
go test ./...
```

## Swagger
### Gerar documentação
```
swag init
```

### Acessar Swagger
```
http://localhost:8080/swagger/index.html
```

Arquivos gerados:
```
/docs
  docs.go
  swagger.json
  swagger.yaml
```

## Como Executar Localmente
### Clone do projeto
```
git clone https://github.com/caioleone/go-grimoire-api.git
```

### Entre na pasta
```
cd go-grimoire-api
```

### Instale dependencias
```
go mod tidy
```

### Execute o projeto
```
go run .
```

Servidor 
```
http://localhost:8080
```

## Ferramentas para testar a API 
* Usuario pode testar usando:
  * Insomnia
  * Postman
  * Curl

## Observação:
O projeto utiliza SQLite para fins educacionais e demonstração.
Em ambiente de produção, recomenda-se PostgreSQL ou MySQL.

## "Banco de Dados" em Memória

Inicialmente, utilizaremos um armazenamento em memória com map.
```
type ID string

type Spell struct {
	ID          ID
	Name        string
	Description string
	ManaCost    int
	Element     string
  Type        string
  Power       int
}

type Creature struct {
	ID          ID
	Name        string
	Description string
	HP          int
	Attack      int
	Defense     int
}

type Application struct {
	spells    map[ID]Spell
	creatures map[ID]Creature
}
```

## Operações do Repositório
* Magias
  1. FindAll()
  2. FindSpellByID(id)
  3. FindByElement()
  4. FindWithFilters()
  5. InsertSpell(spell)
  6. UpdateSpell(id, spell)
  7. DeleteSpell(id)

* Criaturas
  1. FindAllCreatures()
  2. FindCreatureByID(id)
  3. FindWithFilters()
  4. TeachSpells()
  5. InsertCreature(creature)
  6. UpdateCreature(id, creature)
  7. DeleteCreature(id)

## Estrutura do Projeto (Clean Architecture - Simplificado)
```
/cmd
  main.go

/api
  /domain
    spell.go
    creature.go

  /usecase
    spellUsecase.go
    creatureUsecase.go

  /repository
    interfaces.go
    repositoryCreature.go
    repositorySpell.go
    sqliteCreatureRepository.go
    sqliteSpellRepository.go

  /handler
    spellStruct.go
    spellHandler.go
    creatureStruct.go
    creatureHandler.go
```
## Status do projeto

- [X] CRUD Magias
- [X] CRUD Criaturas
- [X] Relacionamento (Creature -> Spell)
- [X] SQLite
- [X] Filtros
- [X] Ordenação
- [X] Paginação
- [X] Clean Architecture básica
- [X] Swagger
- [X] Testes Automatizados
- [X] Docker

## Próximos Desafios (Upgrade do Projeto)

Se quiser evoluir depois:

- [ ] Autenticação JWT
- [ ] Cache com Redis
- [ ] Deploy em produção
- [ ] CI/CD
- [ ] Rate Limiting
- [ ] Logs estruturados

## Objetivo do Projeto
Este projeto foi desenvolvido como parte da jornada de aprendizado em backend com Go, com foco na construção de APIs RESTful organizadas, testáveis e escaláveis utilizando boas práticas de arquitetura e separação de responsabilidades.