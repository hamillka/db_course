import { useLocation, useNavigate } from "react-router-dom";

const Admin: React.FC = () => {
  const navigate = useNavigate();
  const location = useLocation();

  const handleAddAppointmentClick = async () => {
    navigate("/api/admin/appointment_add", {
      state: {
        login: location.state.login,
        password: location.state.password,
        token: location.state.token,
        uid: location.state.uid,
      },
    });
  };

  const handleCancelAppointmentClick = async () => {
    navigate("/api/admin/appointment_delete", {
      state: {
        login: location.state.login,
        password: location.state.password,
        token: location.state.token,
        uid: location.state.uid,
      },
    });
  };

  const handleEditAppointmentClick = async () => {
    navigate("/api/admin/appointment_edit", {
      state: {
        login: location.state.login,
        password: location.state.password,
        token: location.state.token,
        uid: location.state.uid,
      },
    });
  };

  const handleGetAppointmentClick = async () => {
    navigate("/api/admin/appointment_get", {
      state: {
        login: location.state.login,
        password: location.state.password,
        token: location.state.token,
        uid: location.state.uid,
      },
    });
  };

  const handleGetAllDoctorsClick = async () => {
    navigate("/api/admin/doctors_get", {
      state: {
        login: location.state.login,
        password: location.state.password,
        token: location.state.token,
        uid: location.state.uid,
      },
    });
  };

  const handleGetAllPatientsClick = async () => {
    navigate("/api/admin/patients_get", {
      state: {
        login: location.state.login,
        password: location.state.password,
        token: location.state.token,
        uid: location.state.uid,
      },
    });
  };

  return (
    <div>
      <center>
        <button
          style={{
            marginTop: window.innerHeight * 0.4,
            borderRadius: 7.5,
            fontSize: 16,
            marginLeft: 7,
            width: 180,
            height: 50,
            backgroundColor: "cyan",
          }}
          onClick={handleAddAppointmentClick}
        >
          Добавить запись
        </button>

        <button
          style={{
            borderRadius: 7.5,
            fontSize: 16,
            marginLeft: 7,
            marginTop: 15,
            width: 180,
            height: 50,
            backgroundColor: "cyan",
          }}
          onClick={handleCancelAppointmentClick}
        >
          Удалить запись
        </button>

        <button
          style={{
            borderRadius: 7.5,
            fontSize: 16,
            marginLeft: 7,
            marginTop: 15,
            width: 180,
            height: 50,
            backgroundColor: "cyan",
          }}
          onClick={handleEditAppointmentClick}
        >
          Изменить запись
        </button>

        <div>
          <button
            style={{
              borderRadius: 7.5,
              fontSize: 16,
              marginLeft: 7,
              marginTop: 15,
              width: 180,
              height: 50,
              backgroundColor: "cyan",
            }}
            onClick={handleGetAppointmentClick}
          >
            Получить
            <br /> запись
          </button>

          <button
            style={{
              borderRadius: 7.5,
              fontSize: 16,
              marginLeft: 7,
              marginTop: 15,
              width: 180,
              height: 50,
              backgroundColor: "cyan",
            }}
            onClick={handleGetAllDoctorsClick}
          >
            Получить всех врачей
          </button>

          <button
            style={{
              borderRadius: 7.5,
              fontSize: 16,
              marginLeft: 7,
              marginTop: 15,
              width: 180,
              height: 50,
              backgroundColor: "cyan",
            }}
            onClick={handleGetAllPatientsClick}
          >
            Получить всех пациентов
          </button>
        </div>
      </center>
    </div>
  );
};

export default Admin;
