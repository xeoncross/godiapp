package ledis

import (
	"encoding/binary"

	"github.com/Xeoncross/godiapp"
	lediscfg "github.com/siddontang/ledisdb/config"
	"github.com/siddontang/ledisdb/ledis"
)

type DB struct {
	db *ledis.DB
}

func NewDB() (*DB, error) {
	// Use Ledis's default config
	cfg := lediscfg.NewConfigDefault()

	// Put all this in a temp directory
	// cfg.DataDir, err = ioutil.TempDir(os.TempDir(), "")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	l, err := ledis.Open(cfg)
	if err != nil {
		return nil, err
	}

	db, err := l.Select(0)
	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

func (db *DB) GetUsers() (users []*godiapp.User, err error) {

	var pairs []ledis.ScorePair
	pairs, err = db.db.ZRange([]byte("users"), 0, 10)

	var hash []ledis.FVPair
	for _, v := range pairs {
		hash, err = db.db.HGetAll(v.Member)
		if err != nil {
			return
		}

		var user godiapp.User
		for _, v := range hash {
			if string(v.Field) == "ID" {
				user.ID = ByteToInt64(v.Value)
			} else if string(v.Field) == "Email" {
				user.Email = string(v.Value)
			}
		}

		users = append(users, &user)
	}

	return
}

func ByteToInt64(b []byte) int64 {
	x := binary.LittleEndian.Uint64(b)
	return int64(x)
}
