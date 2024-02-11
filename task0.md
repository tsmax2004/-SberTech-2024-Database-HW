# Задание 0
#### Согласно теореме CAP к какой части вы можете отнести следующие СУБД:
- #### DragonFly 
- #### ScyllaDB
- #### ArenadataDB


### 1. DragonFly - CA
[DragonFly](https://www.dragonflydb.io/) - in-memory база данных, то есть она размещена в оперативной памяти. Поэтому нет смысла говорить о `partition tolerance`. Остальные свойства поддерживаются, и, более того, эта база данных специализируется на **высокой доступности**

### 2. ScyllaDB - AP
[ScyllaDB](https://www.scylladb.com/) позволяет извлечь максимум производительности из распределенной системы, гарантируя свойства `availability` и `partition tolerance`, и пренебрегая `consistency` ([документация](https://opensource.docs.scylladb.com/stable/architecture/architecture-fault-tolerance.html#scylladb-architecture-fault-tolerance))

### 3. ArenadataDB - CP
[ArenadataDB](https://arenadata.tech/products/arenadata-db/) - распределенная СУБД, использующая концепцию MPP (massively parallel processing, массивно-параллельные вычисления). Она специализируется на аналитических вычислений, поэтому здесь важна `partition tolerance`. Также ArenadataDB
соответствует принципам ACID, что говорит о `consistency`
