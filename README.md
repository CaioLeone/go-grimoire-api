# API REST - Grimório de Magias e Criaturas
## Conheça o Projeto

Este projeto é um desafio prático em Go para construção de uma API RESTful com temática de fantasia. O objetivo é desenvolver um sistema CRUD para gerenciar Magias e Criaturas, aplicando conceitos fundamentais de HTTP, JSON, organização de projeto e princípios básicos de Clean Architecture.

## Objetivos de Aprendizado
* Trabalhar com net/http
* Manipulação de JSON
* Estruturação de projetos em Go
* Aplicação de Clean Architecture (básico)
* Criação de middlewares
* Persistência com SQLite (extra)

## Domínio da Aplicação

A API será responsável por gerenciar:

* Magias
* Criaturas

## Estrutura das Entidades
1. Magia (Spell)
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

2. Criatura (Creature)
```
{
  "id": "", // UUID, obrigatório
  "name": "Dragão Vermelho",
  "description": "Uma criatura lendária que cospe fogo",
  "hp": 500,
  "attack": 80,
  "defense": 60
}
```
* Regras:
  1. name: obrigatório (mín. 3, máx. 50)
  2. description: obrigatório (mín. 10, máx. 300)
  3. hp, attack, defense: obrigatórios (> 0)

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

## Operações do Repositório
* Magias
  1. FindAllSpells()
  2. FindSpellByID(id)
  3. InsertSpell(spell)
  4. UpdateSpell(id, spell)
  5. DeleteSpell(id)

* Criaturas
  1. FindAllCreatures()
  2. FindCreatureByID(id)
  3. InsertCreature(creature)
  4. UpdateCreature(id, creature)
  5. DeleteCreature(id)

## Endpoints da API
### Magias
  
```POST /api/spells```
Cria uma nova magia

1. 201: retorna magia criada
2. 400: dados inválidos
3. 500: erro interno

```GET /api/spells```
Lista todas as magias

1. 200: lista de magias
2. 500

```GET /api/spells/:id```
Busca uma magia por ID

1. 200
2. 404
3. 500

```PUT /api/spells/:id```
Atualiza uma magia

1. 200
2. 400
3. 404
4. 500

```DELETE /api/spells/:id```
Remove uma magia

1. 200
2. 404
3. 500

### Criaturas
```POST /api/creatures```
Cria uma criatura

```GET /api/creatures```
Lista todas

```GET /api/creatures/:id```
Busca por ID

```PUT /api/creatures/:id```
Atualiza

```DELETE /api/creatures/:id```
Remove

(Mesmas regras de status das magias)

## Padrão de Erro

Todas as respostas de erro devem seguir:
```
{
  "error": "mensagem descritiva"
}
```

## Estrutura do Projeto (Clean Architecture - Simplificado)
```
/cmd
  /api
    main.go

/internal
  /domain
    spell.go
    creature.go

  /usecase
    spell_usecase.go
    creature_usecase.go

  /repository
    memory_repository.go
    sqlite_repository.go (extra)

  /handler
    spell_handler.go
    creature_handler.go

  /middleware
    logging.go
    recovery.go
```
## Middlewares

Você deve implementar:

1. Logging (log de requisições)
2. Recovery (evitar crash da aplicação)
3. Opcional: autenticação simples (API Key)

## Persistência com SQLite (Extra)

* Substituir o armazenamento em memória por SQLite usando:
  1. database/sql
  2. Driver: github.com/mattn/go-sqlite3

* Objetivo:
  1. Criar tabelas spells e creatures
  2. Implementar repositório persistente

### Testes

* Teste a API com:
  1. Postman
  2. Insomnia
  3. Curl

### Próximos Desafios (Upgrade do Projeto)

Se quiser evoluir depois:

Filtros (ex: magias por elemento)
Relacionamento (criatura possui magias)
Paginação
Autenticação JWT
Deploy (Railway, Render)