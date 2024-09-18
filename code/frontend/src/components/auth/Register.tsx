import React, { useState } from "react";
import axios from "axios";
import { useNavigate } from "react-router-dom";
import { ROLE, MenuItem } from "../../types/mytypes";
import DropdownMenu from "../utils/menu";

const Register: React.FC = () => {
  const [fullName, setFullName] = useState("");
  const [phoneNumber, setPhoneNumber] = useState("");
  const [email, setEmail] = useState("");
  const [insurance, setInsurance] = useState("");
  const [specialization, setSpecialization] = useState("");
  const [login, setLogin] = useState("");
  const [password, setPassword] = useState("");
  const [role, setSelectedRole] = useState<MenuItem | null>(null);
  const navigate = useNavigate();

  const handleSelect = (item: MenuItem) => {
    setSelectedRole(item);
  };

  const handleRegisterClick = async () => {
    try {
      await axios.post("http://localhost:8080/auth/register", {
        fio: fullName,
        phoneNumber: phoneNumber,
        email: email,
        insurance: insurance,
        specialization: specialization,
        login: login,
        password: password,
        role: role?.value,
      });
      console.log("Register ok");
      navigate("/");
    } catch (error) {
      alert("Ошибка регистрации");
      console.error("There was an error!", error);
    }
  };

  return (
    <div>
      <center>
        <input
          style={{ marginTop: window.innerHeight * 0.3 }}
          type="text"
          placeholder="ФИО"
          onChange={(e) => setFullName(e.target.value)}
        />
        <input
          type="text"
          placeholder="Номер телефона"
          onChange={(e) => setPhoneNumber(e.target.value)}
        />
        <input
          type="text"
          placeholder="Email"
          onChange={(e) => setEmail(e.target.value)}
        />
        <input
          type="text"
          placeholder="Страховка"
          onChange={(e) => setInsurance(e.target.value)}
        />
        <input
          type="text"
          placeholder="Специализация"
          onChange={(e) => setSpecialization(e.target.value)}
        />
        <div>
          <input
            style={{ marginTop: 10 }}
            type="text"
            placeholder="Логин"
            onChange={(e) => setLogin(e.target.value)}
          />
          <input
            type="password"
            placeholder="Пароль"
            onChange={(e) => setPassword(e.target.value)}
          />
          <DropdownMenu
            items={[
              { label: "Пациент", value: ROLE.USER },
              { label: "Доктор", value: ROLE.DOCTOR },
            ]}
            onSelect={handleSelect}
          />
          <button
            style={{
              borderRadius: 7.5,
              fontSize: 16,
              marginLeft: 7,
              marginTop: 15,
              width: 180,
              backgroundColor: "cyan",
            }}
            onClick={handleRegisterClick}
          >
            Зарегистрироваться
          </button>
        </div>
      </center>
    </div>
  );
};

export default Register;
