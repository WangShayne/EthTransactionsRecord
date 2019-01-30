
CREATE TABLE `ethereum_transtions` (
    `id` int(64) unsigned NOT NULL AUTO_INCREMENT,
    `timestamp` varchar(30) NOT NULL COMMENT '时间戳',
    `from_address` varchar(255) NOT NULL COMMENT '',
    `to_address` varchar(255) NOT NULL COMMENT 'hash值',
    `value_total` varchar(255) NOT NULL COMMENT 'hash值',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='以太坊区块交易记录';
