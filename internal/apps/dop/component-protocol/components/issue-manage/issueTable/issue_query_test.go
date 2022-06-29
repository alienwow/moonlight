// Code generated by MockGen. DO NOT EDIT.
// Source: /Users/fanshuo/Documents/GitHub/erda/internal/apps/dop/providers/issue/core/query/provider.go

// Package issueTable is a generated GoMock package.
package issueTable

import (
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	pb "github.com/erda-project/erda-proto-go/common/pb"
	pb0 "github.com/erda-project/erda-proto-go/dop/issue/core/pb"
	pb1 "github.com/erda-project/erda-proto-go/dop/issue/sync/pb"
	apistructs "github.com/erda-project/erda/apistructs"
	query "github.com/erda-project/erda/internal/apps/dop/providers/issue/core/query"
	dao "github.com/erda-project/erda/internal/apps/dop/providers/issue/dao"
)

// MockInterface is a mock of Interface interface.
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface.
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance.
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// AfterIssueAppRelationCreate mocks base method.
func (m *MockInterface) AfterIssueAppRelationCreate(issueIDs []int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AfterIssueAppRelationCreate", issueIDs)
	ret0, _ := ret[0].(error)
	return ret0
}

// AfterIssueAppRelationCreate indicates an expected call of AfterIssueAppRelationCreate.
func (mr *MockInterfaceMockRecorder) AfterIssueAppRelationCreate(issueIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AfterIssueAppRelationCreate", reflect.TypeOf((*MockInterface)(nil).AfterIssueAppRelationCreate), issueIDs)
}

// AfterIssueInclusionRelationChange mocks base method.
func (m *MockInterface) AfterIssueInclusionRelationChange(id uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AfterIssueInclusionRelationChange", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// AfterIssueInclusionRelationChange indicates an expected call of AfterIssueInclusionRelationChange.
func (mr *MockInterfaceMockRecorder) AfterIssueInclusionRelationChange(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AfterIssueInclusionRelationChange", reflect.TypeOf((*MockInterface)(nil).AfterIssueInclusionRelationChange), id)
}

// AfterIssueUpdate mocks base method.
func (m *MockInterface) AfterIssueUpdate(u *query.IssueUpdated) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AfterIssueUpdate", u)
	ret0, _ := ret[0].(error)
	return ret0
}

// AfterIssueUpdate indicates an expected call of AfterIssueUpdate.
func (mr *MockInterfaceMockRecorder) AfterIssueUpdate(u interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AfterIssueUpdate", reflect.TypeOf((*MockInterface)(nil).AfterIssueUpdate), u)
}

// BatchUpdateIssue mocks base method.
func (m *MockInterface) BatchUpdateIssue(req *pb0.BatchUpdateIssueRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchUpdateIssue", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// BatchUpdateIssue indicates an expected call of BatchUpdateIssue.
func (mr *MockInterfaceMockRecorder) BatchUpdateIssue(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchUpdateIssue", reflect.TypeOf((*MockInterface)(nil).BatchUpdateIssue), req)
}

// CreatePropertyRelation mocks base method.
func (m *MockInterface) CreatePropertyRelation(req *pb0.CreateIssuePropertyInstanceRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePropertyRelation", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreatePropertyRelation indicates an expected call of CreatePropertyRelation.
func (mr *MockInterfaceMockRecorder) CreatePropertyRelation(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePropertyRelation", reflect.TypeOf((*MockInterface)(nil).CreatePropertyRelation), req)
}

// ExportExcel mocks base method.
func (m *MockInterface) ExportExcel(issues []*pb0.Issue, properties []*pb0.IssuePropertyIndex, projectID uint64, isDownload bool, orgID int64, locale string) (io.Reader, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExportExcel", issues, properties, projectID, isDownload, orgID, locale)
	ret0, _ := ret[0].(io.Reader)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ExportExcel indicates an expected call of ExportExcel.
func (mr *MockInterfaceMockRecorder) ExportExcel(issues, properties, projectID, isDownload, orgID, locale interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExportExcel", reflect.TypeOf((*MockInterface)(nil).ExportExcel), issues, properties, projectID, isDownload, orgID, locale)
}

// GetAllIssuesByProject mocks base method.
func (m *MockInterface) GetAllIssuesByProject(req pb0.IssueListRequest) ([]dao.IssueItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllIssuesByProject", req)
	ret0, _ := ret[0].([]dao.IssueItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllIssuesByProject indicates an expected call of GetAllIssuesByProject.
func (mr *MockInterfaceMockRecorder) GetAllIssuesByProject(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllIssuesByProject", reflect.TypeOf((*MockInterface)(nil).GetAllIssuesByProject), req)
}

// GetBatchProperties mocks base method.
func (m *MockInterface) GetBatchProperties(orgID int64, issuesType []string) ([]*pb0.IssuePropertyIndex, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBatchProperties", orgID, issuesType)
	ret0, _ := ret[0].([]*pb0.IssuePropertyIndex)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBatchProperties indicates an expected call of GetBatchProperties.
func (mr *MockInterfaceMockRecorder) GetBatchProperties(orgID, issuesType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBatchProperties", reflect.TypeOf((*MockInterface)(nil).GetBatchProperties), orgID, issuesType)
}

// GetIssue mocks base method.
func (m *MockInterface) GetIssue(id int64, identityInfo *pb.IdentityInfo) (*pb0.Issue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIssue", id, identityInfo)
	ret0, _ := ret[0].(*pb0.Issue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIssue indicates an expected call of GetIssue.
func (mr *MockInterfaceMockRecorder) GetIssue(id, identityInfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIssue", reflect.TypeOf((*MockInterface)(nil).GetIssue), id, identityInfo)
}

// GetIssueChildren mocks base method.
func (m *MockInterface) GetIssueChildren(id uint64, req pb0.PagingIssueRequest) ([]dao.IssueItem, uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIssueChildren", id, req)
	ret0, _ := ret[0].([]dao.IssueItem)
	ret1, _ := ret[1].(uint64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetIssueChildren indicates an expected call of GetIssueChildren.
func (mr *MockInterfaceMockRecorder) GetIssueChildren(id, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIssueChildren", reflect.TypeOf((*MockInterface)(nil).GetIssueChildren), id, req)
}

// GetIssueItem mocks base method.
func (m *MockInterface) GetIssueItem(id uint64) (*dao.IssueItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIssueItem", id)
	ret0, _ := ret[0].(*dao.IssueItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIssueItem indicates an expected call of GetIssueItem.
func (mr *MockInterfaceMockRecorder) GetIssueItem(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIssueItem", reflect.TypeOf((*MockInterface)(nil).GetIssueItem), id)
}

// GetIssueLabelsByProjectID mocks base method.
func (m *MockInterface) GetIssueLabelsByProjectID(projectID uint64) ([]dao.IssueLabel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIssueLabelsByProjectID", projectID)
	ret0, _ := ret[0].([]dao.IssueLabel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIssueLabelsByProjectID indicates an expected call of GetIssueLabelsByProjectID.
func (mr *MockInterfaceMockRecorder) GetIssueLabelsByProjectID(projectID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIssueLabelsByProjectID", reflect.TypeOf((*MockInterface)(nil).GetIssueLabelsByProjectID), projectID)
}

// GetIssueParents mocks base method.
func (m *MockInterface) GetIssueParents(issueID uint64, relationType []string) ([]dao.IssueItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIssueParents", issueID, relationType)
	ret0, _ := ret[0].([]dao.IssueItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIssueParents indicates an expected call of GetIssueParents.
func (mr *MockInterfaceMockRecorder) GetIssueParents(issueID, relationType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIssueParents", reflect.TypeOf((*MockInterface)(nil).GetIssueParents), issueID, relationType)
}

// GetIssueRelationsByIssueIDs mocks base method.
func (m *MockInterface) GetIssueRelationsByIssueIDs(issueID uint64, relationType []string) ([]uint64, []uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIssueRelationsByIssueIDs", issueID, relationType)
	ret0, _ := ret[0].([]uint64)
	ret1, _ := ret[1].([]uint64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetIssueRelationsByIssueIDs indicates an expected call of GetIssueRelationsByIssueIDs.
func (mr *MockInterfaceMockRecorder) GetIssueRelationsByIssueIDs(issueID, relationType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIssueRelationsByIssueIDs", reflect.TypeOf((*MockInterface)(nil).GetIssueRelationsByIssueIDs), issueID, relationType)
}

// GetIssueStage mocks base method.
func (m *MockInterface) GetIssueStage(req *pb0.IssueStageRequest) ([]*pb0.IssueStage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIssueStage", req)
	ret0, _ := ret[0].([]*pb0.IssueStage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIssueStage indicates an expected call of GetIssueStage.
func (mr *MockInterfaceMockRecorder) GetIssueStage(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIssueStage", reflect.TypeOf((*MockInterface)(nil).GetIssueStage), req)
}

// GetIssueStateIDs mocks base method.
func (m *MockInterface) GetIssueStateIDs(req *pb0.GetIssueStatesRequest) ([]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIssueStateIDs", req)
	ret0, _ := ret[0].([]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIssueStateIDs indicates an expected call of GetIssueStateIDs.
func (mr *MockInterfaceMockRecorder) GetIssueStateIDs(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIssueStateIDs", reflect.TypeOf((*MockInterface)(nil).GetIssueStateIDs), req)
}

// GetIssueStateIDsByTypes mocks base method.
func (m *MockInterface) GetIssueStateIDsByTypes(req *apistructs.IssueStatesRequest) ([]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIssueStateIDsByTypes", req)
	ret0, _ := ret[0].([]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIssueStateIDsByTypes indicates an expected call of GetIssueStateIDsByTypes.
func (mr *MockInterfaceMockRecorder) GetIssueStateIDsByTypes(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIssueStateIDsByTypes", reflect.TypeOf((*MockInterface)(nil).GetIssueStateIDsByTypes), req)
}

// GetIssueStatesBelong mocks base method.
func (m *MockInterface) GetIssueStatesBelong(req *pb0.GetIssueStateRelationRequest) ([]apistructs.IssueStateState, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIssueStatesBelong", req)
	ret0, _ := ret[0].([]apistructs.IssueStateState)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIssueStatesBelong indicates an expected call of GetIssueStatesBelong.
func (mr *MockInterfaceMockRecorder) GetIssueStatesBelong(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIssueStatesBelong", reflect.TypeOf((*MockInterface)(nil).GetIssueStatesBelong), req)
}

// GetIssueStatesMap mocks base method.
func (m *MockInterface) GetIssueStatesMap(req *pb0.GetIssueStatesRequest) (map[string][]pb0.IssueStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIssueStatesMap", req)
	ret0, _ := ret[0].(map[string][]pb0.IssueStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIssueStatesMap indicates an expected call of GetIssueStatesMap.
func (mr *MockInterfaceMockRecorder) GetIssueStatesMap(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIssueStatesMap", reflect.TypeOf((*MockInterface)(nil).GetIssueStatesMap), req)
}

// GetIssuesByIssueIDs mocks base method.
func (m *MockInterface) GetIssuesByIssueIDs(issueIDs []uint64) ([]*pb0.Issue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIssuesByIssueIDs", issueIDs)
	ret0, _ := ret[0].([]*pb0.Issue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIssuesByIssueIDs indicates an expected call of GetIssuesByIssueIDs.
func (mr *MockInterfaceMockRecorder) GetIssuesByIssueIDs(issueIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIssuesByIssueIDs", reflect.TypeOf((*MockInterface)(nil).GetIssuesByIssueIDs), issueIDs)
}

// GetIssuesStatesByProjectID mocks base method.
func (m *MockInterface) GetIssuesStatesByProjectID(projectID uint64, issueType string) ([]dao.IssueState, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIssuesStatesByProjectID", projectID, issueType)
	ret0, _ := ret[0].([]dao.IssueState)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIssuesStatesByProjectID indicates an expected call of GetIssuesStatesByProjectID.
func (mr *MockInterfaceMockRecorder) GetIssuesStatesByProjectID(projectID, issueType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIssuesStatesByProjectID", reflect.TypeOf((*MockInterface)(nil).GetIssuesStatesByProjectID), projectID, issueType)
}

// GetProperties mocks base method.
func (m *MockInterface) GetProperties(req *pb0.GetIssuePropertyRequest) ([]*pb0.IssuePropertyIndex, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProperties", req)
	ret0, _ := ret[0].([]*pb0.IssuePropertyIndex)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProperties indicates an expected call of GetProperties.
func (mr *MockInterfaceMockRecorder) GetProperties(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProperties", reflect.TypeOf((*MockInterface)(nil).GetProperties), req)
}

// ListStatesTransByProjectID mocks base method.
func (m *MockInterface) ListStatesTransByProjectID(projectID uint64) ([]dao.IssueStateTransition, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListStatesTransByProjectID", projectID)
	ret0, _ := ret[0].([]dao.IssueStateTransition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListStatesTransByProjectID indicates an expected call of ListStatesTransByProjectID.
func (mr *MockInterfaceMockRecorder) ListStatesTransByProjectID(projectID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListStatesTransByProjectID", reflect.TypeOf((*MockInterface)(nil).ListStatesTransByProjectID), projectID)
}

// Paging mocks base method.
func (m *MockInterface) Paging(req pb0.PagingIssueRequest) ([]*pb0.Issue, uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Paging", req)
	ret0, _ := ret[0].([]*pb0.Issue)
	ret1, _ := ret[1].(uint64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Paging indicates an expected call of Paging.
func (mr *MockInterfaceMockRecorder) Paging(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Paging", reflect.TypeOf((*MockInterface)(nil).Paging), req)
}

// SyncIssueChildrenIteration mocks base method.
func (m *MockInterface) SyncIssueChildrenIteration(issue *pb0.Issue, iterationID int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncIssueChildrenIteration", issue, iterationID)
	ret0, _ := ret[0].(error)
	return ret0
}

// SyncIssueChildrenIteration indicates an expected call of SyncIssueChildrenIteration.
func (mr *MockInterfaceMockRecorder) SyncIssueChildrenIteration(issue, iterationID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncIssueChildrenIteration", reflect.TypeOf((*MockInterface)(nil).SyncIssueChildrenIteration), issue, iterationID)
}

// SyncLabels mocks base method.
func (m *MockInterface) SyncLabels(value *pb1.Value, issueIDs []uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncLabels", value, issueIDs)
	ret0, _ := ret[0].(error)
	return ret0
}

// SyncLabels indicates an expected call of SyncLabels.
func (mr *MockInterfaceMockRecorder) SyncLabels(value, issueIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncLabels", reflect.TypeOf((*MockInterface)(nil).SyncLabels), value, issueIDs)
}

// UpdateIssue mocks base method.
func (m *MockInterface) UpdateIssue(req *pb0.UpdateIssueRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateIssue", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateIssue indicates an expected call of UpdateIssue.
func (mr *MockInterfaceMockRecorder) UpdateIssue(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateIssue", reflect.TypeOf((*MockInterface)(nil).UpdateIssue), req)
}

// UpdateLabels mocks base method.
func (m *MockInterface) UpdateLabels(id, projectID uint64, labelNames []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateLabels", id, projectID, labelNames)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateLabels indicates an expected call of UpdateLabels.
func (mr *MockInterfaceMockRecorder) UpdateLabels(id, projectID, labelNames interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateLabels", reflect.TypeOf((*MockInterface)(nil).UpdateLabels), id, projectID, labelNames)
}