package config

import (
	"fmt"
	"time"

	"github.com/Le-BlitzZz/apartments-rabbit-app/dataserver/internal/entity"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	dbServer   = "mariadb:4001"
	dbName     = "apartments"
	dbUser     = "apartments"
	dbPassword = "apartments"
	dbTimeout  = 5
)

func (c *Config) Db() *gorm.DB {
	if c.db == nil {
		log.Fatal("config: database not connected")
	}

	return c.db
}

func (c *Config) DatabaseUser() string {
	if c.options.DatabaseUser == "" {
		return dbUser
	}
	return c.options.DatabaseUser
}

func (c *Config) DatabasePassword() string {
	if c.options.DatabasePassword == "" {
		return dbPassword
	}
	return c.options.DatabasePassword
}

func (c *Config) DatabaseServer() string {
	if c.options.DatabaseServer == "" {
		return dbServer
	}
	return c.options.DatabaseServer
}

func (c *Config) DatabaseName() string {
	if c.options.DatabaseName == "" {
		return dbName
	}
	return c.options.DatabaseName
}

func (c *Config) DatabaseTimeout() int {
	if c.options.DatabaseTimeout == 0 {
		return dbTimeout
	}
	return c.options.DatabaseTimeout
}

func (c *Config) DatabaseDsn() string {
	dbServer := fmt.Sprintf("tcp(%s)", c.DatabaseServer())

	return fmt.Sprintf(
		"%s:%s@%s/%s?charset=utf8mb4,utf8&collation=utf8mb4_unicode_ci&parseTime=true&timeout=%ds",
		c.DatabaseUser(),
		c.DatabasePassword(),
		dbServer,
		c.DatabaseName(),
		c.DatabaseTimeout(),
	)
}

func (c *Config) InitDb() {
	entity.SetDb(c.db)
	c.migrateDb()
}

func (c *Config) connectDb() error {
	dbDsn := c.DatabaseDsn()
	gormCfg := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	}

	db, err := gorm.Open(mysql.Open(dbDsn), gormCfg)

	if err != nil || db == nil {
		log.Println("MariaDB: waiting to become available")

		for i := 1; i <= 12; i++ {
			db, err = gorm.Open(mysql.Open(dbDsn), gormCfg)

			if db != nil && err == nil {
				break
			}

			time.Sleep(5 * time.Second)
		}

		if err != nil || db == nil {
			return err
		}
	}

	log.Println("MariaDB: connection established")

	c.db = db

	return nil
}

func (c *Config) closeDb() error {
	if c.db != nil {
		sqlDB, err := c.db.DB()
		if err != nil {
			return err
		}
		if err := sqlDB.Close(); err != nil {
			return err
		}

		c.db = nil
	}

	return nil
}

func (c *Config) migrateDb() {
	for name, entity := range entity.Entities {
		if err := c.db.AutoMigrate(entity); err != nil {
			log.Errorf("Failed migrating %s", name)
		}
	}
}
