# API REST - Grimório de Magias e Criaturas
## Conheça o Projeto

Este projeto é um desafio prático em Go para construção de uma API RESTful com temática de fantasia, focado no gerenciamento de ***Magias*** e ***Criaturas***. 

O objetivo é evoluir de um CRUD simples para uma arquitetura mais robusta, aplicando conceitos como: 
* Clean Architecture (básico)
* Separação de responsabilidade
* Persistencia com SQLite
* Relacionamentos enre entidades
* Filtros, ordenação e paginação

## Objetivos de Aprendizado
* Trabalhar com ```net/http```
* Manipulação de JSON
* Uso de router ```chi```
* Estruturação de projetos em Go
* Aplicação de Clean Architecture
* Uso de interfaces
* Criação de middlewares
* Persistência com SQLite 
* Queries dinâmicas (filtros)
* Paginação e ordenação

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
}
```

* Regras:

  1. name: obrigatório (mín. 3, máx. 50)
  2. description: obrigatório (mín. 10, máx. 300)
  3. mana_cost: obrigatório (> 0)
  4. element: obrigatório (ex: Fogo, Gelo, Arcano)

2. ***Criatura (Creature)***
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
  1. name: obrigatório (mín. 3, máx. 50)
  2. description: obrigatório (mín. 10, máx. 300)
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

## Padrão de Erro

Todas as respostas de erro devem seguir:
```
{
  "error": "mensagem descritiva"
}
```

## Middlewares
* Logging
* Recovery
* RequestID

## Testes
* Usuario pode testar usando:
  * Insomnia
  * Postman
  * Curl

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

## Próximos Desafios (Upgrade do Projeto)

Se quiser evoluir depois:

- [ ] Swagger
- [ ] Autenticação
- [ ] Cache
- [ ] Testes Automatizados
- [ ] Deploy (Vercel)