import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import Admin from "./components/ActionsAdmin";
import Doctor from "./components/ActionsDoctor";
import Patient from "./components/ActionsPatient";
import Login from "./components/auth/Login";
import Register from "./components/auth/Register";
import AppointmentAdd from "./components/admin/AppointmentAdd";
import AppointmentDelete from "./components/admin/AppointmentDelete";
import AppointmentEdit from "./components/admin/AppointmentEdit";
import AppointmentGet from "./components/admin/AppointmentGet";
import DoctorsGet from "./components/admin/DoctorsGet";
import PatientsGet from "./components/admin/PatientsGet";
import MyAppointmentDelete from "./components/patient/AppointmentDelete";
import MyAppointmentAdd from "./components/patient/AppointmentAdd";
import MyAppointmentsGet from "./components/patient/AppointmentsGet";
import MyAppointmentEdit from "./components/patient/AppointmentEdit";
import GetMyMedHistory from "./components/patient/MedHistoryGet";

import AppointmentsGet from "./components/doctor/AppointmentsGet";
import CreateMedHistory from "./components/doctor/CreateMedHistory";
import GetMedHistory from "./components/doctor/GetMedHistoryByPatient";
import UpdateMedHistory from "./components/doctor/UpdateMedHistoryByPatient";

function App() {
  return (
    <Router>
      <div>
        <Routes>
          <Route path="/" element={<Login />} />
          <Route path="/register" element={<Register />} />

          <Route path="/api/admin" element={<Admin />} />
          <Route
            path="/api/admin/appointment_get"
            element={<AppointmentGet />}
          />
          <Route
            path="/api/admin/appointment_add"
            element={<AppointmentAdd />}
          />
          <Route
            path="/api/admin/appointment_edit"
            element={<AppointmentEdit />}
          />
          <Route
            path="/api/admin/appointment_delete"
            element={<AppointmentDelete />}
          />
          <Route path="/api/admin/doctors_get" element={<DoctorsGet />} />
          <Route path="/api/admin/patients_get" element={<PatientsGet />} />

          <Route path="/api/patient" element={<Patient />} />
          <Route
            path="/api/patient/appointment_delete"
            element={<MyAppointmentDelete />}
          />
          <Route
            path="/api/patient/appointment_add"
            element={<MyAppointmentAdd />}
          />
          <Route
            path="/api/patient/appointments_get"
            element={<MyAppointmentsGet />}
          />
          <Route
            path="/api/patient/appointment_edit"
            element={<MyAppointmentEdit />}
          />
          <Route
            path="/api/patient/medhistory_get"
            element={<GetMyMedHistory />}
          />

          <Route path="/api/doctor" element={<Doctor />} />

          <Route
            path="/api/doctor/appointment_get"
            element={<AppointmentsGet />}
          />
          <Route
            path="/api/doctor/medhistory_get"
            element={<GetMedHistory />}
          />
          <Route
            path="/api/doctor/medhistory_create"
            element={<CreateMedHistory />}
          />
          <Route
            path="/api/doctor/medhistory_update"
            element={<UpdateMedHistory />}
          />
        </Routes>
      </div>
    </Router>
  );
}

export default App;
