import React, { useState } from "react";
import axios from "axios";
import { useLocation, useNavigate } from "react-router-dom";

const AppointmentDelete: React.FC = () => {
  const [id, setId] = useState(0);

  const location = useLocation();
  const navigate = useNavigate();

  let headers = {
    "Content-Type": "application/json",
    "auth-x": `Bearer ${location.state.token}`,
  };

  const handleDelClick = async () => {
    try {
      await axios.delete(`http://localhost:8080/api/v1/appointment/${id}`, {
        headers: headers,
      });
      navigate("/api/admin", {
        state: {
          login: location.state.login,
          password: location.state.password,
          token: location.state.token,
          uid: location.state.uid,
        },
      });
    } catch (error) {
      alert("Ошибка удаления записи");
      console.error("There was an error!", error);
    }
  };

  return (
    <div>
      <center>
        <input
          style={{ marginTop: window.innerHeight * 0.3, height: 30 }}
          type="number"
          placeholder="ID записи"
          onChange={(e) => setId(parseInt(e.target.value))}
        />
        <button
          style={{
            borderRadius: 7.5,
            fontSize: 16,
            marginLeft: 7,
            marginTop: 15,
            width: 180,
            height: 30,
            backgroundColor: "cyan",
          }}
          onClick={handleDelClick}
        >
          Удалить
        </button>
      </center>
    </div>
  );
};

export default AppointmentDelete;
