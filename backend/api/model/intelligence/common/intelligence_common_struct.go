// Code generated by thriftgo (0.4.1). DO NOT EDIT.

package common

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
)

type IntelligenceStatus int64

const (
	IntelligenceStatus_Using   IntelligenceStatus = 1
	IntelligenceStatus_Deleted IntelligenceStatus = 2
	IntelligenceStatus_Banned  IntelligenceStatus = 3
	// 迁移失败
	IntelligenceStatus_MoveFailed IntelligenceStatus = 4
	// 复制中
	IntelligenceStatus_Copying IntelligenceStatus = 5
	// 复制失败
	IntelligenceStatus_CopyFailed IntelligenceStatus = 6
)

func (p IntelligenceStatus) String() string {
	switch p {
	case IntelligenceStatus_Using:
		return "Using"
	case IntelligenceStatus_Deleted:
		return "Deleted"
	case IntelligenceStatus_Banned:
		return "Banned"
	case IntelligenceStatus_MoveFailed:
		return "MoveFailed"
	case IntelligenceStatus_Copying:
		return "Copying"
	case IntelligenceStatus_CopyFailed:
		return "CopyFailed"
	}
	return "<UNSET>"
}

func IntelligenceStatusFromString(s string) (IntelligenceStatus, error) {
	switch s {
	case "Using":
		return IntelligenceStatus_Using, nil
	case "Deleted":
		return IntelligenceStatus_Deleted, nil
	case "Banned":
		return IntelligenceStatus_Banned, nil
	case "MoveFailed":
		return IntelligenceStatus_MoveFailed, nil
	case "Copying":
		return IntelligenceStatus_Copying, nil
	case "CopyFailed":
		return IntelligenceStatus_CopyFailed, nil
	}
	return IntelligenceStatus(0), fmt.Errorf("not a valid IntelligenceStatus string")
}

func IntelligenceStatusPtr(v IntelligenceStatus) *IntelligenceStatus { return &v }
func (p *IntelligenceStatus) Scan(value interface{}) (err error) {
	var result sql.NullInt64
	err = result.Scan(value)
	*p = IntelligenceStatus(result.Int64)
	return
}

func (p *IntelligenceStatus) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	return int64(*p), nil
}

type IntelligenceType int64

const (
	IntelligenceType_Bot     IntelligenceType = 1
	IntelligenceType_Project IntelligenceType = 2
)

func (p IntelligenceType) String() string {
	switch p {
	case IntelligenceType_Bot:
		return "Bot"
	case IntelligenceType_Project:
		return "Project"
	}
	return "<UNSET>"
}

func IntelligenceTypeFromString(s string) (IntelligenceType, error) {
	switch s {
	case "Bot":
		return IntelligenceType_Bot, nil
	case "Project":
		return IntelligenceType_Project, nil
	}
	return IntelligenceType(0), fmt.Errorf("not a valid IntelligenceType string")
}

func IntelligenceTypePtr(v IntelligenceType) *IntelligenceType { return &v }
func (p *IntelligenceType) Scan(value interface{}) (err error) {
	var result sql.NullInt64
	err = result.Scan(value)
	*p = IntelligenceType(result.Int64)
	return
}

func (p *IntelligenceType) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	return int64(*p), nil
}

type IntelligenceBasicInfo struct {
	ID             int64              `thrift:"id,1" form:"id" json:"id,string" query:"id"`
	Name           string             `thrift:"name,2" form:"name" json:"name" query:"name"`
	Description    string             `thrift:"description,3" form:"description" json:"description" query:"description"`
	IconURI        string             `thrift:"icon_uri,4" form:"icon_uri" json:"icon_uri" query:"icon_uri"`
	IconURL        string             `thrift:"icon_url,5" form:"icon_url" json:"icon_url" query:"icon_url"`
	SpaceID        int64              `thrift:"space_id,6" form:"space_id" json:"space_id,string" query:"space_id"`
	OwnerID        int64              `thrift:"owner_id,7" form:"owner_id" json:"owner_id,string" query:"owner_id"`
	CreateTime     int64              `thrift:"create_time,8" form:"create_time" json:"create_time,string" query:"create_time"`
	UpdateTime     int64              `thrift:"update_time,9" form:"update_time" json:"update_time,string" query:"update_time"`
	Status         IntelligenceStatus `thrift:"status,10" form:"status" json:"status" query:"status"`
	PublishTime    int64              `thrift:"publish_time,11" form:"publish_time" json:"publish_time,string" query:"publish_time"`
	EnterpriseID   *string            `thrift:"enterprise_id,12,optional" form:"enterprise_id" json:"enterprise_id,omitempty" query:"enterprise_id"`
	OrganizationID *int64             `thrift:"organization_id,13,optional" form:"organization_id" json:"organization_id,omitempty" query:"organization_id"`
}

func NewIntelligenceBasicInfo() *IntelligenceBasicInfo {
	return &IntelligenceBasicInfo{}
}

func (p *IntelligenceBasicInfo) InitDefault() {
}

func (p *IntelligenceBasicInfo) GetID() (v int64) {
	return p.ID
}

func (p *IntelligenceBasicInfo) GetName() (v string) {
	return p.Name
}

func (p *IntelligenceBasicInfo) GetDescription() (v string) {
	return p.Description
}

func (p *IntelligenceBasicInfo) GetIconURI() (v string) {
	return p.IconURI
}

func (p *IntelligenceBasicInfo) GetIconURL() (v string) {
	return p.IconURL
}

func (p *IntelligenceBasicInfo) GetSpaceID() (v int64) {
	return p.SpaceID
}

func (p *IntelligenceBasicInfo) GetOwnerID() (v int64) {
	return p.OwnerID
}

func (p *IntelligenceBasicInfo) GetCreateTime() (v int64) {
	return p.CreateTime
}

func (p *IntelligenceBasicInfo) GetUpdateTime() (v int64) {
	return p.UpdateTime
}

func (p *IntelligenceBasicInfo) GetStatus() (v IntelligenceStatus) {
	return p.Status
}

func (p *IntelligenceBasicInfo) GetPublishTime() (v int64) {
	return p.PublishTime
}

var IntelligenceBasicInfo_EnterpriseID_DEFAULT string

func (p *IntelligenceBasicInfo) GetEnterpriseID() (v string) {
	if !p.IsSetEnterpriseID() {
		return IntelligenceBasicInfo_EnterpriseID_DEFAULT
	}
	return *p.EnterpriseID
}

var IntelligenceBasicInfo_OrganizationID_DEFAULT int64

func (p *IntelligenceBasicInfo) GetOrganizationID() (v int64) {
	if !p.IsSetOrganizationID() {
		return IntelligenceBasicInfo_OrganizationID_DEFAULT
	}
	return *p.OrganizationID
}

var fieldIDToName_IntelligenceBasicInfo = map[int16]string{
	1:  "id",
	2:  "name",
	3:  "description",
	4:  "icon_uri",
	5:  "icon_url",
	6:  "space_id",
	7:  "owner_id",
	8:  "create_time",
	9:  "update_time",
	10: "status",
	11: "publish_time",
	12: "enterprise_id",
	13: "organization_id",
}

func (p *IntelligenceBasicInfo) IsSetEnterpriseID() bool {
	return p.EnterpriseID != nil
}

func (p *IntelligenceBasicInfo) IsSetOrganizationID() bool {
	return p.OrganizationID != nil
}

func (p *IntelligenceBasicInfo) Read(iprot thrift.TProtocol) (err error) {
	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I64 {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 3:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField3(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 4:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField4(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 5:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField5(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 6:
			if fieldTypeId == thrift.I64 {
				if err = p.ReadField6(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 7:
			if fieldTypeId == thrift.I64 {
				if err = p.ReadField7(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 8:
			if fieldTypeId == thrift.I64 {
				if err = p.ReadField8(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 9:
			if fieldTypeId == thrift.I64 {
				if err = p.ReadField9(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 10:
			if fieldTypeId == thrift.I32 {
				if err = p.ReadField10(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 11:
			if fieldTypeId == thrift.I64 {
				if err = p.ReadField11(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 12:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField12(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 13:
			if fieldTypeId == thrift.I64 {
				if err = p.ReadField13(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_IntelligenceBasicInfo[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *IntelligenceBasicInfo) ReadField1(iprot thrift.TProtocol) error {

	var _field int64
	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		_field = v
	}
	p.ID = _field
	return nil
}
func (p *IntelligenceBasicInfo) ReadField2(iprot thrift.TProtocol) error {

	var _field string
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		_field = v
	}
	p.Name = _field
	return nil
}
func (p *IntelligenceBasicInfo) ReadField3(iprot thrift.TProtocol) error {

	var _field string
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		_field = v
	}
	p.Description = _field
	return nil
}
func (p *IntelligenceBasicInfo) ReadField4(iprot thrift.TProtocol) error {

	var _field string
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		_field = v
	}
	p.IconURI = _field
	return nil
}
func (p *IntelligenceBasicInfo) ReadField5(iprot thrift.TProtocol) error {

	var _field string
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		_field = v
	}
	p.IconURL = _field
	return nil
}
func (p *IntelligenceBasicInfo) ReadField6(iprot thrift.TProtocol) error {

	var _field int64
	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		_field = v
	}
	p.SpaceID = _field
	return nil
}
func (p *IntelligenceBasicInfo) ReadField7(iprot thrift.TProtocol) error {

	var _field int64
	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		_field = v
	}
	p.OwnerID = _field
	return nil
}
func (p *IntelligenceBasicInfo) ReadField8(iprot thrift.TProtocol) error {

	var _field int64
	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		_field = v
	}
	p.CreateTime = _field
	return nil
}
func (p *IntelligenceBasicInfo) ReadField9(iprot thrift.TProtocol) error {

	var _field int64
	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		_field = v
	}
	p.UpdateTime = _field
	return nil
}
func (p *IntelligenceBasicInfo) ReadField10(iprot thrift.TProtocol) error {

	var _field IntelligenceStatus
	if v, err := iprot.ReadI32(); err != nil {
		return err
	} else {
		_field = IntelligenceStatus(v)
	}
	p.Status = _field
	return nil
}
func (p *IntelligenceBasicInfo) ReadField11(iprot thrift.TProtocol) error {

	var _field int64
	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		_field = v
	}
	p.PublishTime = _field
	return nil
}
func (p *IntelligenceBasicInfo) ReadField12(iprot thrift.TProtocol) error {

	var _field *string
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		_field = &v
	}
	p.EnterpriseID = _field
	return nil
}
func (p *IntelligenceBasicInfo) ReadField13(iprot thrift.TProtocol) error {

	var _field *int64
	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		_field = &v
	}
	p.OrganizationID = _field
	return nil
}

func (p *IntelligenceBasicInfo) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("IntelligenceBasicInfo"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}
		if err = p.writeField3(oprot); err != nil {
			fieldId = 3
			goto WriteFieldError
		}
		if err = p.writeField4(oprot); err != nil {
			fieldId = 4
			goto WriteFieldError
		}
		if err = p.writeField5(oprot); err != nil {
			fieldId = 5
			goto WriteFieldError
		}
		if err = p.writeField6(oprot); err != nil {
			fieldId = 6
			goto WriteFieldError
		}
		if err = p.writeField7(oprot); err != nil {
			fieldId = 7
			goto WriteFieldError
		}
		if err = p.writeField8(oprot); err != nil {
			fieldId = 8
			goto WriteFieldError
		}
		if err = p.writeField9(oprot); err != nil {
			fieldId = 9
			goto WriteFieldError
		}
		if err = p.writeField10(oprot); err != nil {
			fieldId = 10
			goto WriteFieldError
		}
		if err = p.writeField11(oprot); err != nil {
			fieldId = 11
			goto WriteFieldError
		}
		if err = p.writeField12(oprot); err != nil {
			fieldId = 12
			goto WriteFieldError
		}
		if err = p.writeField13(oprot); err != nil {
			fieldId = 13
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *IntelligenceBasicInfo) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("id", thrift.I64, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI64(p.ID); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}
func (p *IntelligenceBasicInfo) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("name", thrift.STRING, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Name); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}
func (p *IntelligenceBasicInfo) writeField3(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("description", thrift.STRING, 3); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Description); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 end error: ", p), err)
}
func (p *IntelligenceBasicInfo) writeField4(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("icon_uri", thrift.STRING, 4); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.IconURI); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 4 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 4 end error: ", p), err)
}
func (p *IntelligenceBasicInfo) writeField5(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("icon_url", thrift.STRING, 5); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.IconURL); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 5 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 5 end error: ", p), err)
}
func (p *IntelligenceBasicInfo) writeField6(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("space_id", thrift.I64, 6); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI64(p.SpaceID); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 6 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 6 end error: ", p), err)
}
func (p *IntelligenceBasicInfo) writeField7(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("owner_id", thrift.I64, 7); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI64(p.OwnerID); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 7 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 7 end error: ", p), err)
}
func (p *IntelligenceBasicInfo) writeField8(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("create_time", thrift.I64, 8); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI64(p.CreateTime); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 8 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 8 end error: ", p), err)
}
func (p *IntelligenceBasicInfo) writeField9(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("update_time", thrift.I64, 9); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI64(p.UpdateTime); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 9 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 9 end error: ", p), err)
}
func (p *IntelligenceBasicInfo) writeField10(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("status", thrift.I32, 10); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI32(int32(p.Status)); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 10 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 10 end error: ", p), err)
}
func (p *IntelligenceBasicInfo) writeField11(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("publish_time", thrift.I64, 11); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI64(p.PublishTime); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 11 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 11 end error: ", p), err)
}
func (p *IntelligenceBasicInfo) writeField12(oprot thrift.TProtocol) (err error) {
	if p.IsSetEnterpriseID() {
		if err = oprot.WriteFieldBegin("enterprise_id", thrift.STRING, 12); err != nil {
			goto WriteFieldBeginError
		}
		if err := oprot.WriteString(*p.EnterpriseID); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 12 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 12 end error: ", p), err)
}
func (p *IntelligenceBasicInfo) writeField13(oprot thrift.TProtocol) (err error) {
	if p.IsSetOrganizationID() {
		if err = oprot.WriteFieldBegin("organization_id", thrift.I64, 13); err != nil {
			goto WriteFieldBeginError
		}
		if err := oprot.WriteI64(*p.OrganizationID); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 13 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 13 end error: ", p), err)
}

func (p *IntelligenceBasicInfo) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("IntelligenceBasicInfo(%+v)", *p)

}
