import React, { useState } from "react";
import axios from "axios";
import { useLocation } from "react-router-dom";

interface Patient {
  fio: string;
  phoneNumber: string;
  email: string;
  insurance: string;
  id: number;
}

const PatientsGet: React.FC = () => {
  const location = useLocation();

  const [patients, setPatients] = useState<Patient[]>();

  let headers = {
    "Content-Type": "application/json",
    "auth-x": `Bearer ${location.state.token}`,
  };

  const handleGetClick = async () => {
    try {
      const response = await axios.get<Patient[]>(
        `http://localhost:8080/api/v1/patient`,
        {
          headers: headers,
        },
      );
      setPatients(response.data);
    } catch (error) {
      alert("Ошибка получения списка пациентов");
      console.error("There was an error!", error);
    }
  };

  return (
    <div>
      <center>
        <button
          style={{
            borderRadius: 7.5,
            fontSize: 16,
            marginLeft: 7,
            marginTop: 15,
            width: 180,
            backgroundColor: "cyan",
          }}
          onClick={handleGetClick}
        >
          Вывести
        </button>

        <h2>Пациенты</h2>
        <table style={{ borderCollapse: "collapse", width: "100%" }}>
          <thead>
            <tr style={{ borderBottom: "1px solid #ddd" }}>
              <th style={{ textAlign: "center", padding: "8px" }}>
                ФИО пациента
              </th>
              <th style={{ textAlign: "center", padding: "8px" }}>Почта</th>
              <th style={{ textAlign: "center", padding: "8px" }}>
                Номер телефона
              </th>
              <th style={{ textAlign: "center", padding: "8px" }}>Страховка</th>
            </tr>
          </thead>
          <tbody>
            <tr style={{ borderBottom: "1px solid #ddd" }}>
              <td style={{ textAlign: "center", padding: "8px" }}>
                {patients?.map((item) => <div key={item.id}> {item.fio}</div>)}
              </td>
              <td style={{ textAlign: "center", padding: "8px" }}>
                {patients?.map((item) => <div key={item.id}> {item.email}</div>)}{" "}
              </td>
              <td style={{ textAlign: "center", padding: "8px" }}>
                {patients?.map((item) => (
                  <div key={item.id}> {item.phoneNumber}</div>
                ))}
              </td>
              <td style={{ textAlign: "center", padding: "8px" }}>
                {patients?.map((item) => (
                  <div key={item.id}> {item.insurance}</div>
                ))}
              </td>
            </tr>
          </tbody>
        </table>
      </center>
    </div>
  );
};

export default PatientsGet;
