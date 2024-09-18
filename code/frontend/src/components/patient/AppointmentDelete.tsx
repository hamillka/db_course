import React, { useState, useEffect, useMemo } from "react";
import axios from "axios";
import { useLocation, useNavigate } from "react-router-dom";

interface Item {
  id: number;
  dateTime: string;
  patientId: number;
  doctorId: number;
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

const MyAppointmentDelete: React.FC = () => {
  const [data, setData] = useState<Item[]>([]);
  const [item, setItem] = useState<Item | null>(null);

  const location = useLocation();
  const navigate = useNavigate();

  const headers = useMemo(() => {
    return {
      "Content-Type": "application/json",
      "auth-x": `Bearer ${location.state.token}`,
    };
  }, [location.state.token]);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await axios.get<Item[]>(
          `http://localhost:8080/api/v1/appointment?patient_id=${location.state.pid}`,
          {
            headers: headers,
          },
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
        setData(appointmentsWithDoctors);
      } catch (error) {
        console.error("Error fetching data:", error);
      }
    };

    fetchData();
  }, [location, headers]);

  const handleItemClick = (item: Item) => {
    setItem(item);
  };

  const handleDelClick = async () => {
    try {
      await axios.delete(
        `http://localhost:8080/api/v1/appointment/${item?.id}`,
        {
          headers: headers,
        },
      );
      navigate("/api/patient", {
        state: {
          login: location.state.login,
          password: location.state.password,
          token: location.state.token,
          uid: location.state.uid,
          pid: location.state.pid,
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
        <h1>Ваши записи</h1>
        <ul>
          {data.map((item) => (
            <div key={item.id} onClick={() => handleItemClick(item)}>
              Дата и время:{" "}
                  {item?.dateTime
                    ? new Date(item?.dateTime).toLocaleString("ru-RU", {
                        year: "numeric",
                        month: "long",
                        day: "numeric",
                        hour: "2-digit",
                        minute: "2-digit",
                        timeZone: "UTC",
                      })
                    : ""}
            </div>
          ))}
        </ul>
        {item && (
          <div>
            <h2>Выбранная запись</h2>
            <table style={{ borderCollapse: "collapse", width: "100%" }}>
              <thead>
                <tr style={{ borderBottom: "1px solid #ddd" }}>
                  <th style={{ textAlign: "center", padding: "8px" }}>
                    ФИО врача
                  </th>
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
                    {item?.docName}
                  </td>
                  <td style={{ textAlign: "center", padding: "8px" }}>
                    {item?.docSpec}
                  </td>
                  <td style={{ textAlign: "center", padding: "8px" }}>
                    {item?.dateTime
                      ? new Date(item?.dateTime).toLocaleString("ru-RU", {
                          year: "numeric",
                          month: "long",
                          day: "numeric",
                          hour: "2-digit",
                          minute: "2-digit",
                          timeZone: "UTC",
                        })
                      : ""}
                  </td>
                </tr>
              </tbody>
            </table>
            <button
              style={{
                borderRadius: 7.5,
                fontSize: 16,
                marginLeft: 7,
                marginTop: 15,
                width: 180,
                backgroundColor: "cyan",
              }}
              onClick={handleDelClick}
            >
              Отменить выбранную запись
            </button>
          </div>
        )}
      </center>
    </div>
  );
};

export default MyAppointmentDelete;
