const asyncHandler = require("express-async-handler");
const { getAllStudents, addNewStudent, getStudentDetail, setStudentStatus, updateStudent } = require("./students-service");

const handleGetAllStudents = asyncHandler(async (req, res) => {
    const payload = req.query;
    const students = await getAllStudents(payload);
    res.status(200).json({ data: students });

});

const handleAddStudent = asyncHandler(async (req, res) => {
    const payload = req.body;
    const result = await addNewStudent(payload);
    res.status(201).json({ message: result.message });

});

const handleUpdateStudent = asyncHandler(async (req, res) => {
    const payload = { ...req.body, id: req.params.id };
    const result = await updateStudent(payload);
    res.status(200).json({ message: result.message });

});

const handleGetStudentDetail = asyncHandler(async (req, res) => {
    const studentId = req.params.id;
    const student = await getStudentDetail(studentId);
    res.status(200).json({ data: student });

});

const handleStudentStatus = asyncHandler(async (req, res) => {
    const { id: userId } = req.params;
    const { reviewerId, status } = req.body;
    const result = await setStudentStatus({ userId, reviewerId, status });
    res.status(200).json({ message: result.message });

});

module.exports = {
    handleGetAllStudents,
    handleGetStudentDetail,
    handleAddStudent,
    handleStudentStatus,
    handleUpdateStudent,
};
