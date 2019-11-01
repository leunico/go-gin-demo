package models

type ExamineeDeviceProbling struct {
    Model

    ExaminationExamineeId uint `json:"examination_examinee_id"`
    IsCamera int8 `json:"is_camera"`
    IsMicrophone int8 `json:"is_microphone"`
    IsChrome int8 `json:"is_chrome"`
    IsMcIde int8 `json:"is_mc_ide"`
    IsScratchIde int8 `json:"is_scratch_ide"`
    IsPythonIde int8 `json:"is_python_ide"`
    IsCIde int8 `json:"is_c_ide"`
}

func (ExamineeDeviceProbling) TableName() string {
    return "examinee_device_probings"
}