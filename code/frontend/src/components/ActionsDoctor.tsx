import { useLocation, useNavigate } from "react-router-dom";

const Doctor: React.FC = () => {
  const navigate = useNavigate();
  const location = useLocation();

  const handleGetAppointmentsClick = async () => {
    navigate("/api/doctor/appointment_get", {
      state: {
        login: location.state.login,
        password: location.state.password,
        token: location.state.token,
        uid: location.state.uid,
        did: location.state.did,
      },
    });
  };

  const handleGetMedHistoryByPatientIdClick = async () => {
    navigate("/api/doctor/medhistory_get", {
      state: {
        login: location.state.login,
        password: location.state.password,
        token: location.state.token,
        uid: location.state.uid,
        did: location.state.did,
      },
    });
  };

  const handleCreateMedHistoryClick = async () => {
    navigate("/api/doctor/medhistory_create", {
      state: {
        login: location.state.login,
        password: location.state.password,
        token: location.state.token,
        uid: location.state.uid,
        did: location.state.did,
      },
    });
  };

  const handleUpdateMedHistoryByPatientClick = async () => {
    navigate("/api/doctor/medhistory_update", {
      state: {
        login: location.state.login,
        password: location.state.password,
        token: location.state.token,
        uid: location.state.uid,
        did: location.state.did,
      },
    });
  };

  return (
    <div>
      <center>
        <button
          style={{
            marginTop: window.innerHeight * 0.2,
            borderRadius: 7.5,
            fontSize: 16,
            marginLeft: 7,
            width: 180,
            height: 60,
            backgroundColor: "cyan",
          }}
          onClick={handleGetAppointmentsClick}
        >
          Получить <br />
          записанных ко мне <br /> пациентов
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
          onClick={handleGetMedHistoryByPatientIdClick}
        >
          Получить
          <br /> медицинскую карту
          <br /> пациента
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
          onClick={handleCreateMedHistoryClick}
        >
          Создать
          <br /> медицинскую карту
          <br /> пациента
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
          onClick={handleUpdateMedHistoryByPatientClick}
        >
          Изменить
          <br /> медицинскую карту
          <br /> пациента
        </button>
      </center>
    </div>
  );
};

export default Doctor;
