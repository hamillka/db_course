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

const GetMyMedHistory: React.FC = () => {
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
        `http://localhost:8080/api/v1/medical_history/${location.state.pid}`,
        { headers: headers },
      );
      setMedHistory(response.data);
    } catch (error) {
      alert("Ошибка получения мед. карты");
      console.error("There was and error!", error);
    }
  };

  return (
    <div>
      <center>
        <h1> Получение медицинской карты</h1>
        <button
          style={{
            borderRadius: 7.5,
            fontSize: 16,
            marginLeft: 7,
            marginTop: 15,
            marginBottom: 25,
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
              {medHistory?.bloodType}
              </td>
              <td style={{ textAlign: "center", padding: "8px" }}>
                {medHistory?.chronicDiseases}
              </td>
              <td style={{ textAlign: "center", padding: "8px" }}>
                {medHistory?.allergies}
              </td>
              <td style={{ textAlign: "center", padding: "8px" }}>
                {medHistory?.vaccination}
              </td>
            </tr>
          </tbody>
        </table>
      </center>
    </div>
  );
};

export default GetMyMedHistory;
