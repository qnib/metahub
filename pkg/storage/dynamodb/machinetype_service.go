package dynamodb

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/qnib/metahub/pkg/storage"
)

type machineTypeService struct {
	ctx context.Context
}

func formatLogin(accountName string, login string) {

}
func (s *machineTypeService) Init() (err error) {
	return
}

func (s *machineTypeService) GetByID(accountName string, id int64) (mt *storage.MachineType, err error) {
	log.Printf("GetByID(%s, %d)\n", accountName, id)
	return mt, err
}

/*** GetByUsername will be used to authenticate a client.
We'll chop off the prefix (e.g. qnib from qnib-c518xl-shp2) and chat the prefix against the users table
*/
func (s *machineTypeService) GetByUsername(username string) (mt *storage.MachineType, err error) {
	log.Printf("GetByUsername(%s)\n", username)
	// Chop of the first part of the usrename, without the type specific suffixes
	userSplit := strings.Split(username, "-")
	if len(userSplit) == 1 {
		err = fmt.Errorf("username should contain the actual user seperated by a dash (e.g. qnib-type1). Got: %s", username)
		return
	}
	usern := userSplit[0]
	user, err := mhTableUserScan(svc, fmt.Sprintf("%s_users", mhDbTablePrefix), usern)
	if err != nil {
		return
	}
	mt = &storage.MachineType{
		Login:    user.Login,
		Password: user.Password,
	}
	typen := strings.Join(userSplit[1:], "-")
	typ, err := mhTableTypeScan(svc, fmt.Sprintf("%s_types", mhDbTablePrefix), typen)
	if err != nil {
		return
	}
	mt.Features = strings.Split(typ.Features, ",")
	mt.DisplayName = typ.Type
	return
}

func (s *machineTypeService) Add(accountName string, mt *storage.MachineType) (err error) {
	return err
}

func (s *machineTypeService) Delete(accountName string, id int64) error {
	return nil
}

func (s *machineTypeService) List(accountName string) (mts []storage.MachineType, err error) {
	log.Printf("mt.List(accountName=%s)", accountName)
	typeItems, err := mhTableTypeList(svc, fmt.Sprintf("%s_types", mhDbTablePrefix))
	if err != nil {
		return
	}
	for _, tItem := range typeItems {
		mt := storage.MachineType{
			DisplayName: tItem.Type,
			Features:    strings.Split(tItem.Features, ","),
		}
		mts = append(mts, mt)
	}
	return
}

func (s *machineTypeService) Update(accountName string, mt storage.MachineType) (err error) {
	return err
}
