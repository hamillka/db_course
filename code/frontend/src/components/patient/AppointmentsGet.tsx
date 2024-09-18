import React, { useState, useMemo } from "react";
import axios from "axios";
import { useLocation } from "react-router-dom";
import { useEffect } from "react";

interface Appointment {
  id: number;
  patientId: number;
  doctorId: number;
  dateTime: string;
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

const MyAppointmentsGet: React.FC = () => {
  const location = useLocation();

  const [appointments, setAppointments] = useState<Appointment[]>([]);

  const headers = useMemo(() => {
    return {
      "Content-Type": "application/json",
      "auth-x": `Bearer ${location.state.token}`,
    };
  }, [location.state.token]);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await axios.get<Appointment[]>(
          `http://localhost:8080/api/v1/appointment?patient_id=${location.state.pid}`,
          { headers: headers },
        );
        const appointmentsWithDoctors = await Promise.all(
          response.data.map(async (appointment) => {
            const responseDoc = await axios.get<Doctor>(
              `http://localhost:8080/api/v1/doctor/${appointment.doctorId}`,
              { headers: headers },
            );
            return {
              ...appointment,
              docName: responseDoc.data.fio,
              docSpec: responseDoc.data.specialization,
            };
          }),
        );
        setAppointments(appointmentsWithDoctors);
      } catch (error) {
        alert("Ошибка получения записей");
        console.error("Error fetching data:", error);
      }
    };
    fetchData();
  }, [headers, location.state.pid]);

  return (
    <div>
      <center>
        <h2>Ваши записи</h2>
        <table style={{ borderCollapse: "collapse", width: "100%" }}>
          <thead>
            <tr style={{ borderBottom: "1px solid #ddd" }}>
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
                {appointments.map((item) => (
                  <div key={item.id}> {item.docName}</div>
                ))}
              </td>
              <td style={{ textAlign: "center", padding: "8px" }}>
                {appointments.map((item) => (
                  <div key={item.id}> {item.docSpec}</div>
                ))}
              </td>
              <td style={{ textAlign: "center", padding: "8px" }}>
                {appointments.map((item) => (
                  <div key={item.id}> 
                   {item?.dateTime
                      ? new Date(item?.dateTime).toLocaleString("ru-RU", {
                          year: "numeric",
                          month: "long",
                          day: "numeric",
                          hour: "2-digit",
                          minute: "2-digit",
                          timeZone: "UTC"
                        })
                      : ""}
                  </div>
                ))}
              </td>
            </tr>
          </tbody>
        </table>
      </center>
    </div>
  );
};

export default MyAppointmentsGet;
