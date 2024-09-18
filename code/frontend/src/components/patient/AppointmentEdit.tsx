import React, { useState, useMemo, useEffect } from "react";
import axios from "axios";
import { useLocation, useNavigate } from "react-router-dom";
import DatePicker from "react-datepicker";
import "react-datepicker/dist/react-datepicker.css";

interface Doctor {
  fio: string;
  phoneNumber: string;
  email: string;
  id: number;
  specialization: string;
}

interface TimeTable {
  id: number;
  doctorID: number;
  officeID: number;
  workDay: number;
}

interface Appointment {
  id: number;
  patientId: number;
  doctorId: number;
  dateTime: string;
}

const MyAppointmentEdit: React.FC = () => {
  const [data, setData] = useState<Doctor[]>([]);
  const [appData, setAppData] = useState<Appointment[]>([]);
  const [doctor, setDoctor] = useState<Doctor | null>(null);

  const [myDoctor, setMyDoctor] = useState<Doctor | null>(null);

  const [appointment, setAppointment] = useState<Appointment | null>(null);
  const [date, setDate] = useState<Date | null>(null);
  const [tt, setTt] = useState<TimeTable[]>();
  const [okDays, setOkDays] = useState<number[]>([]);

  const location = useLocation();
  const navigate = useNavigate();

  const headers = useMemo(() => {
    return {
      "Content-Type": "application/json",
      "auth-x": `Bearer ${location.state.token}`,
    };
  }, [location.state.token]);

  function dateToRFC3339WithOffset(date: Date, hoursOffset: number): string {
    const offsetMilliseconds = hoursOffset * 60 * 60 * 1000;
    const dateWithOffset = new Date(date.getTime() + offsetMilliseconds);
    let rfc3339String = dateWithOffset.toISOString().slice(0, -5);
    const offsetPrefix = hoursOffset >= 0 ? "+" : "-";
    const offsetHours = Math.abs(hoursOffset).toString().padStart(2, "0");
    const offsetMinutes = "00";
    rfc3339String += `${offsetPrefix}${offsetHours}:${offsetMinutes}`;

    return rfc3339String;
  }

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await axios.get(
          `http://localhost:8080/api/v1/doctor`,
          {
            headers: headers,
          },
        );
        setData(response.data);
      } catch (error) {
        alert("Ошибка изменения записи");
        console.error("Error fetching data:", error);
      }
    };

    fetchData();
  }, [headers, doctor]);

  const getWorkDays = async (doctor: Doctor | null) => {
    try {
      if (doctor?.id) {
        const response = await axios.get<TimeTable[]>(
          `http://localhost:8080/api/v1/workdays/${doctor?.id}`,
          {
            headers: headers,
          },
        );
        return response.data;
      }
    } catch (error) {
      console.error("Error fetching data:", error);
    }
  };

  useEffect(() => {
    const fetchData = async () => {
      const result = await getWorkDays(doctor);
      setTt(result);
    };

    fetchData();
  }, [doctor]);

  useEffect(() => {
    let onlyDays: number[] = [];
    if (tt) {
      for (let i = 0; i < tt.length; i++) {
        onlyDays.push(tt[i].workDay);
      }
      setOkDays(onlyDays);
    }
  }, [tt]);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await axios.get(
          `http://localhost:8080/api/v1/appointment?patient_id=${location.state.pid}`,
          {
            headers: headers,
          },
        );
        setAppData(response.data);
      } catch (error) {
        console.error("Error fetching data:", error);
      }
    };

    fetchData();
  }, [location.state.pid, headers]);

  const handleDoctorClick = async (doctor: Doctor) => {
    setDoctor(doctor);
  };

  const handleAppointmentClick = async (appointment: Appointment) => {
    setAppointment(appointment);
    const response = await axios.get<Doctor>(
      `http://localhost:8080/api/v1/doctor/${appointment.doctorId}`,
      {
        headers: headers,
      },
    );
    setMyDoctor(response.data);
  };

  const handleEditClick = async () => {
    try {
      const rfc3339 = dateToRFC3339WithOffset(date!, 3);
      await axios.patch(
        `http://localhost:8080/api/v1/appointment/${appointment?.id}`,
        {
          patientId: location.state.pid,
          doctorId: doctor?.id,
          dateTime: rfc3339,
        },
        { headers: headers },
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
      console.error("There was an error!", error);
    }
  };

  const isWorkDay = (date: Date) => {
    const day = date.getDay();
    return okDays.indexOf(day) > -1;
  };

  return (
    <div>
      <center>
        <h2>Ваши записи</h2>
        <ul>
          {appData.map((item) => (
            <div key={item.id} onClick={() => handleAppointmentClick(item)}>
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
        <h2>Наши врачи</h2>
        <ul>
          {data.map((item) => (
            <div key={item.id} onClick={() => handleDoctorClick(item)}>
              ФИО: {item.fio}
            </div>
          ))}
        </ul>
        {doctor && appointment && (
          <div>
            <h2>Изменяемая запись</h2>
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
                    {myDoctor?.fio}
                  </td>
                  <td style={{ textAlign: "center", padding: "8px" }}>
                    {myDoctor?.specialization}
                  </td>
                  <td style={{ textAlign: "center", padding: "8px" }}>
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
                  </td>
                </tr>
              </tbody>
            </table>

            <h2>Выбранный врач</h2>
            <table style={{ borderCollapse: "collapse", width: "100%" }}>
              <thead>
                <tr style={{ borderBottom: "1px solid #ddd" }}>
                  <th style={{ textAlign: "center", padding: "8px" }}>ФИО</th>
                  <th style={{ textAlign: "center", padding: "8px" }}>
                    Специализация
                  </th>

                  <th style={{ textAlign: "center", padding: "8px" }}>
                    Номер телефона
                  </th>
                  <th style={{ textAlign: "center", padding: "8px" }}>
                    Дата и время записи
                  </th>
                </tr>
              </thead>
              <tbody>
                <tr style={{ borderBottom: "1px solid #ddd" }}>
                  <td style={{ textAlign: "center", padding: "8px" }}>
                    {doctor.fio}
                  </td>
                  <td style={{ textAlign: "center", padding: "8px" }}>
                    {doctor.specialization}
                  </td>
                  <td style={{ textAlign: "center", padding: "8px" }}>
                    {doctor.phoneNumber}
                  </td>
                  <td style={{ textAlign: "center", padding: "8px" }}>
                    <DatePicker
                      selected={date}
                      onChange={(date) => setDate(date)}
                      showTimeSelect
                      timeFormat="HH:mm"
                      timeIntervals={30}
                      showIcon
                      calendarStartDay={1}
                      timeCaption="Время"
                      dateFormat="MMMM d, yyyy HH:mm"
                      filterDate={isWorkDay}
                      filterTime={(time) => {
                        const selectedHour = time.getHours();
                        return selectedHour >= 8 && selectedHour < 20;
                      }}
                    />
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
                height: 30,
                backgroundColor: "cyan",
              }}
              onClick={handleEditClick}
            >
              Изменить запись
            </button>
          </div>
        )}
      </center>
    </div>
  );
};

export default MyAppointmentEdit;
