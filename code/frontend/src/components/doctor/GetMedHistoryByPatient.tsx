import React, { useState } from "react";
import axios from "axios";
import { useLocation, useNavigate } from "react-router-dom";

interface History {
  chronicDiseases: string;
  allergies: string;
  bloodType: string;
  vaccination: string;
  id: number;
  patientId: number;
}

const GetMedHistory: React.FC = () => {
  const [patientId, setPatientId] = useState<number>(0);
  const [medHistory, setMedHistory] = useState<History>();

  const location = useLocation();
  const navigate = useNavigate();

  const headers = {
    "Content-Type": "application/json",
    "auth-x": `Bearer ${location.state.token}`,
  };

  const handleGetClick = async () => {
    try {
      const response = await axios.get<History>(
        `http://localhost:8080/api/v1/medical_history/${patientId}`,
        { headers: headers },
      );
      setMedHistory(response.data);
      // navigate("/api/doctor", {
      //   state: {
      //     login: location.state.login,
      //     password: location.state.password,
      //     token: location.state.token,
      //     uid: location.state.uid,
      //     did: location.state.pid,
      //   },
      // });
    } catch (error) {
      alert("Ошибка получения мед. карты");
      console.error("There was and error!", error);
    }
  };

  return (
    <div>
      <center>
        <h1> Получение медицинской карты</h1>
        <input
          type="number"
          style={{ height: 30 }}
          placeholder="ID пациента"
          onChange={(e) => setPatientId(parseInt(e.target.value))}
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
          onClick={handleGetClick}
        >
          Получить
        </button>

        <table style={{ borderCollapse: "collapse", width: "100%" }}>
          <thead>
            <tr style={{ borderBottom: "1px solid #ddd" }}>
              <th style={{ textAlign: "center", padding: "8px" }}>
                ID пациента
              </th>
              <th style={{ textAlign: "center", padding: "8px" }}>
                Группа крови
              </th>
              <th style={{ textAlign: "center", padding: "8px" }}>
                Хронические заболевания
              </th>
              <th style={{ textAlign: "center", padding: "8px" }}>Аллергия</th>
              <th style={{ textAlign: "center", padding: "8px" }}>
                Вакцинация
              </th>
            </tr>
          </thead>
          <tbody>
            <tr style={{ borderBottom: "1px solid #ddd" }}>
              <td style={{ textAlign: "center", padding: "8px" }}>
                <div key={medHistory?.id}> {medHistory?.patientId}</div>
              </td>
              <td style={{ textAlign: "center", padding: "8px" }}>
                <div key={medHistory?.id}> {medHistory?.bloodType}</div>
              </td>

              <td style={{ textAlign: "center", padding: "8px" }}>
                <div key={medHistory?.id}> {medHistory?.chronicDiseases}</div>
              </td>
              <td style={{ textAlign: "center", padding: "8px" }}>
                <div key={medHistory?.id}> {medHistory?.allergies}</div>
              </td>
              <td style={{ textAlign: "center", padding: "8px" }}>
                <div key={medHistory?.id}> {medHistory?.vaccination}</div>
              </td>
            </tr>
          </tbody>
        </table>
      </center>
    </div>
  );
};

export default GetMedHistory;
