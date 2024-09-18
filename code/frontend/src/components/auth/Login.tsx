import React, { useState } from "react";
import axios from "axios";
import { useNavigate } from "react-router-dom";
import { jwtDecode } from "jwt-decode";
import { ROLE, MyToken } from "../../types/mytypes";

const Login: React.FC = () => {
  const [login, setLogin] = useState("");
  const [password, setPassword] = useState("");
  const navigate = useNavigate();

  const handleLoginClick = async () => {
    try {
      const response = await axios.post("http://localhost:8080/auth/login", {
        login: login,
        password: password,
      });
      const token = response.data.jwtToken;
      const decoded = jwtDecode<MyToken>(token);
      const user = response.data.user;

      switch (decoded.role) {
        case ROLE.ADMIN: {
          navigate("/api/admin", {
            state: {
              login: login,
              password: password,
              token: token,
              uid: user.id,
            },
          });
          break;
        }
        case ROLE.USER: {
          navigate("/api/patient", {
            state: {
              login: login,
              password: password,
              token: token,
              uid: user.id,
              pid: user.patientId,
            },
          });
          break;
        }
        case ROLE.DOCTOR: {
          navigate("/api/doctor", {
            state: {
              login: login,
              password: password,
              token: token,
              uid: user.id,
              did: user.doctorId,
            },
          });
          break;
        }
      }
    } catch (error) {
      alert("Ошибка входа в систему");
      console.error("There was an error!", error);
    }
  };

  const handleRegisterClick = () => {
    navigate("/register");
  };

  return (
    <div>
      <center>
        <input
          style={{
            marginTop: window.innerHeight * 0.3,
          }}
          type="text"
          placeholder="Введите логин"
          onChange={(e) => setLogin(e.target.value)}
        />
        <input
          style={{
            marginLeft: 10,
          }}
          type="password"
          placeholder="Введите пароль"
          onChange={(e) => setPassword(e.target.value)}
        />
      </center>

      <center>
        <button
          style={{
            width: 180,
            marginTop: 20,
            borderRadius: 7.5,
            backgroundColor: "cyan",
            fontSize: 16,
          }}
          onClick={handleLoginClick}
        >
          Войти
        </button>
      </center>
      <center>
        <button
          style={{
            borderRadius: 7.5,
            marginTop: 10,
            width: 180,
            backgroundColor: "cyan",
            fontSize: 16,
          }}
          onClick={handleRegisterClick}
        >
          Зарегистрироваться
        </button>
      </center>
    </div>
  );
};

export default Login;
