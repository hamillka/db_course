import React, { useState } from "react";
import axios from "axios";
import { useLocation, useNavigate } from "react-router-dom";

const UpdateMedHistory: React.FC = () => {
  const [diseases, setDiseases] = useState<string>("");
  const [allergies, setAllergies] = useState<string>("");
  const [patientId, setPatientId] = useState<number>(0);
  const [bloodType, setBloodType] = useState<string>("");
  const [vaccination, setVaccination] = useState<string>("");

  const location = useLocation();
  const navigate = useNavigate();

  const headers = {
    "Content-Type": "application/json",
    "auth-x": `Bearer ${location.state.token}`,
  };

  const handleUpdateClick = async () => {
    try {
      await axios.patch(
        `http://localhost:8080/api/v1/medical_history/${patientId}`,
        {
          chronicDiseases: diseases,
          allergies: allergies,
          bloodType: bloodType,
          vaccination: vaccination,
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
      alert("Ошибка изменения мед. карты");
      console.error("There was and error!", error);
    }
  };

  return (
    <div>
      <center>
        <h1> Обновление медицинской карты</h1>
        <input
          type="number"
          style={{
            marginTop: window.innerHeight * 0.3,
            height: 30,
            width: 200,
          }}
          placeholder="ID пациента"
          onChange={(e) => setPatientId(parseInt(e.target.value))}
        />
        <input
          type="string"
          style={{ height: 30, width: 200 }}
          placeholder="Группа крови"
          onChange={(e) => setBloodType(e.target.value)}
        />

        <input
          type="string"
          style={{ height: 30, width: 200 }}
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
          onChange={(e) => setVaccination(e.target.value)}
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
          onClick={handleUpdateClick}
        >
          Изменить
        </button>
      </center>
    </div>
  );
};

export default UpdateMedHistory;
