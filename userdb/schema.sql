CREATE TABLE IF NOT EXISTS `tbusers` (
  `id` char(36) NOT NULL,
  `name` varchar(45) NOT NULL,
  `email` varchar(45) NOT NULL,
  `birthdate` date NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idtbusers_UNIQUE` (`id`)
);

CREATE TABLE IF NOT EXISTS `tbproducts` (
  `id` char(36) NOT NULL,
  `descricao` varchar(45) NOT NULL,
  `categoria` varchar(45) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`)
);