import { useLocation, useNavigate } from "react-router-dom";

const Patient: React.FC = () => {
  const navigate = useNavigate();
  const location = useLocation();

  const handleCancelMyAppointmentClick = async () => {
    navigate("/api/patient/appointment_delete", {
      state: {
        login: location.state.login,
        password: location.state.password,
        token: location.state.token,
        uid: location.state.uid,
        pid: location.state.pid,
      },
    });
  };

  const handleEditMyAppointmentClick = async () => {
    navigate("/api/patient/appointment_edit", {
      state: {
        login: location.state.login,
        password: location.state.password,
        token: location.state.token,
        uid: location.state.uid,
        pid: location.state.pid,
      },
    });
  };

  const handleAddMyAppointmentClick = async () => {
    navigate("/api/patient/appointment_add", {
      state: {
        login: location.state.login,
        password: location.state.password,
        token: location.state.token,
        uid: location.state.uid,
        pid: location.state.pid,
      },
    });
  };

  const handleGetMyAppointmentsClick = async () => {
    navigate("/api/patient/appointments_get", {
      state: {
        login: location.state.login,
        password: location.state.password,
        token: location.state.token,
        uid: location.state.uid,
        pid: location.state.pid,
      },
    });
  };

  const handleGetMyMedHistoryClick = async () => {
    navigate("/api/patient/medhistory_get", {
      state: {
        login: location.state.login,
        password: location.state.password,
        token: location.state.token,
        uid: location.state.uid,
        pid: location.state.pid,
      },
    });
  };

  return (
    <div>
      <center>
        <button
          style={{
            marginTop: window.innerHeight * 0.3,
            borderRadius: 7.5,
            fontSize: 16,
            marginLeft: 7,
            width: 180,
            backgroundColor: "cyan",
          }}
          onClick={handleCancelMyAppointmentClick}
        >
          Отменить мою запись
        </button>

        <button
          style={{
            borderRadius: 7.5,
            fontSize: 16,
            marginLeft: 7,
            marginTop: 15,
            width: 180,
            backgroundColor: "cyan",
          }}
          onClick={handleEditMyAppointmentClick}
        >
          Изменить
          <br /> мою запись
        </button>

        <button
          style={{
            borderRadius: 7.5,
            fontSize: 16,
            marginLeft: 7,
            marginTop: 15,
            width: 180,
            backgroundColor: "cyan",
          }}
          onClick={handleAddMyAppointmentClick}
        >
          Записаться
          <br /> ко врачу
        </button>

        <button
          style={{
            borderRadius: 7.5,
            fontSize: 16,
            marginLeft: 7,
            marginTop: 15,
            width: 180,
            backgroundColor: "cyan",
          }}
          onClick={handleGetMyAppointmentsClick}
        >
          Получить все мои записи
        </button>

        <button
          style={{
            borderRadius: 7.5,
            fontSize: 16,
            marginLeft: 7,
            marginTop: 15,
            width: 180,
            backgroundColor: "cyan",
          }}
          onClick={handleGetMyMedHistoryClick}
        >
          Получить мою медицинскую карту
        </button>
      </center>
    </div>
  );
};

export default Patient;
