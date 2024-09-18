import React, { useState } from "react";
import axios from "axios";
import { useLocation, useNavigate } from "react-router-dom";

const CreateMedHistory: React.FC = () => {
  const [diseases, setDiseases] = useState<string>("");
  const [allergies, setAllergies] = useState<string>("");
  const [patientId, setPatientId] = useState<number>(0);
  const [bloodType, setbloodType] = useState<string>("");
  const [vaccination, setvaccination] = useState<string>("");

  const location = useLocation();
  const navigate = useNavigate();

  const headers = {
    "Content-Type": "application/json",
    "auth-x": `Bearer ${location.state.token}`,
  };

  const handleCreateClick = async () => {
    try {
      await axios.post(
        "http://localhost:8080/api/v1/medical_history",
        {
          chronicDiseases: diseases,
          allergies: allergies,
          patientId: patientId,
          vaccination: vaccination,
          bloodType: bloodType,
        },
        { headers: headers },
      );
      navigate("/api/doctor", {
        state: {
          login: location.state.login,
          password: location.state.password,
          token: location.state.token,
          uid: location.state.uid,
          did: location.state.pid,
        },
      });
    } catch (error) {
      alert("Ошибка создания мед. карты");
      console.error("There was and error!", error);
    }
  };

  return (
    <div>
      <center>
        <h1> Создание медицинской карты</h1>
        <input
          type="number"
          style={{ height: 30, width: 200 }}
          placeholder="ID пациента"
          onChange={(e) => setPatientId(parseInt(e.target.value))}
        />
        <input
          style={{
            marginTop: window.innerHeight * 0.3,
            height: 30,
            width: 200,
          }}
          type="string"
          placeholder="Группа крови"
          onChange={(e) => setbloodType(e.target.value)}
        />
        <input
          style={{
            marginTop: window.innerHeight * 0.3,
            height: 30,
            width: 200,
          }}
          type="string"
          placeholder="Хронические заболевания"
          onChange={(e) => setDiseases(e.target.value)}
        />
        <input
          type="string"
          style={{ height: 30, width: 200 }}
          placeholder="Аллергия"
          onChange={(e) => setAllergies(e.target.value)}
        />
        <input
          type="string"
          style={{ height: 30, width: 200 }}
          placeholder="Вакцинация"
          onChange={(e) => setvaccination(e.target.value)}
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
          onClick={handleCreateClick}
        >
          Создать
        </button>
      </center>
    </div>
  );
};

export default CreateMedHistory;
