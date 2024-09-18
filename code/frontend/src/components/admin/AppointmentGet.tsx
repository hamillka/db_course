import React, { useState, useEffect, useMemo } from "react";
import axios from "axios";
import { useLocation } from "react-router-dom";

interface Appointment {
  id: number;
  patientId: number;
  doctorId: number;
  dateTime: string;
  patientName: string;
  docName: string;
  docSpec: string;
}

interface Doctor {
  fio: string;
  phoneNumber: string;
  email: string;
  id: number;
  specialization: string;
}

interface Patient {
  fio: string;
  phoneNumber: string;
  email: string;
  insurance: string;
  id: number;
}

const AppointmentGet: React.FC = () => {
  const [id, setId] = useState(0);

  const location = useLocation();

  const [appointment, setAppointment] = useState<Appointment>({
    id: 0,
    patientId: 0,
    doctorId: 0,
    dateTime: "",
    patientName: "",
    docName: "",
    docSpec: "",
  });

  const headers = useMemo(() => {
    return {
      "Content-Type": "application/json",
      "auth-x": `Bearer ${location.state.token}`,
    };
  }, [location.state.token]);

  const handleGetClick = async () => {
    try {
      const response = await axios.get<Appointment>(
        `http://localhost:8080/api/v1/appointment/${id}`,
        {
          headers: headers,
        },
      );
      setAppointment(response.data);
    } catch (error) {
      alert("Ошибка получения записи");
      console.error("There was an error!", error);
    }
  };

  useEffect(() => {
    const fetchData = async () => {
      try {
        const responseDoc = await axios.get<Doctor>(
          `http://localhost:8080/api/v1/doctor/${appointment.doctorId}`,
          { headers: headers },
        );
        const responsePatient = await axios.get<Patient>(
          `http://localhost:8080/api/v1/patient/${appointment.patientId}`,
          { headers: headers },
        );
        console.log(responseDoc);

        setAppointment({
          id: appointment.id,
          patientId: appointment.patientId,
          doctorId: appointment.doctorId,
          dateTime: appointment.dateTime,
          docName: responseDoc.data.fio,
          docSpec: responseDoc.data.specialization,
          patientName: responsePatient.data.fio,
        });
      } catch (error) {
        console.error("Error fetching data:", error);
      }
    };
    fetchData();
  }, [
    headers,
    appointment.doctorId,
    appointment.patientId,
    appointment.dateTime,
    appointment.id,
  ]);

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
          onClick={handleGetClick}
        >
          Получить
        </button>
        <h2>Запись</h2>
        <table style={{ borderCollapse: "collapse", width: "100%" }}>
          <thead>
            <tr style={{ borderBottom: "1px solid #ddd" }}>
              <th style={{ textAlign: "center", padding: "8px" }}>
                ФИО пациента
              </th>
              <th style={{ textAlign: "center", padding: "8px" }}>ФИО врача</th>
              <th style={{ textAlign: "center", padding: "8px" }}>
                Специализация врача
              </th>
              <th style={{ textAlign: "center", padding: "8px" }}>
                Дата и время записи
              </th>
            </tr>
          </thead>
          <tbody>
            <tr style={{ borderBottom: "1px solid #ddd" }}>
              <td style={{ textAlign: "center", padding: "8px" }}>
                {<div key={appointment?.id}> {appointment?.patientName}</div>}
              </td>
              <td style={{ textAlign: "center", padding: "8px" }}>
                {<div key={appointment?.id}> {appointment?.docName}</div>}
              </td>
              <td style={{ textAlign: "center", padding: "8px" }}>
                {<div key={appointment?.id}> {appointment?.docSpec}</div>}
              </td>

              <td style={{ textAlign: "center", padding: "8px" }}>
                {<div key={appointment?.id}> 
                  {appointment?.dateTime
                    ? new Date(appointment?.dateTime).toLocaleString("ru-RU", {
                        year: "numeric",
                        month: "long",
                        day: "numeric",
                        hour: "2-digit",
                        minute: "2-digit",
                        timeZone: "UTC",
                      })
                    : ""}
                </div>}
              </td>
            </tr>
          </tbody>
        </table>
      </center>
    </div>
  );
};

export default AppointmentGet;
