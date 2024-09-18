import React, { useState, useMemo, useEffect } from "react";
import axios from "axios";
import { useLocation, useNavigate } from "react-router-dom";
import DatePicker from "react-datepicker";
import "react-datepicker/dist/react-datepicker.css";

interface Item {
  fio: string;
  phoneNumber: string;
  email: string;
  id: number;
  specialization: string;
}

interface TimeTable {
  id: number;
  doctorId: number;
  officeId: number;
  workDay: number;
}

const MyAppointmentAdd: React.FC = () => {
  const [data, setData] = useState<Item[]>([]);
  const [item, setItem] = useState<Item | null>(null);

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
        console.error("Error fetching data:", error);
      }
    };

    fetchData();
  }, [headers, item]);

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

  const getWorkDays = async (item: Item | null) => {
    try {
      if (item?.id) {
        const response = await axios.get<TimeTable[]>(
          `http://localhost:8080/api/v1/workdays/${item?.id}`,
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
      const result = await getWorkDays(item);
      setTt(result);
    };

    fetchData();
  }, [item]);

  useEffect(() => {
    let onlyDays: number[] = [];
    if (tt) {
      for (let i = 0; i < tt.length; i++) {
        onlyDays.push(tt[i].workDay);
      }
      setOkDays(onlyDays);
    }
  }, [tt]);

  const handleItemClick = async (item: Item) => {
    setItem(item);
  };

  const handleAddClick = async () => {
    try {
      const rfc3339 = dateToRFC3339WithOffset(date!, +3);

      await axios.post(
        "http://localhost:8080/api/v1/appointment",
        {
          patientId: location.state.pid,
          doctorId: item?.id,
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
      alert("Ошибка создания записи");
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
        <h1>Запись на прием</h1>
        <ul>
          {data.map((item) => (
            <div key={item.id} onClick={() => handleItemClick(item)}>
              ФИО: {item.fio}
            </div>
          ))}
        </ul>
        {item && (
          <div>
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
                    {item.fio}
                  </td>
                  <td style={{ textAlign: "center", padding: "8px" }}>
                    {item.specialization}
                  </td>
                  <td style={{ textAlign: "center", padding: "8px" }}>
                    {item.phoneNumber}
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
              onClick={handleAddClick}
            >
              Подтвердить запись
            </button>
          </div>
        )}
      </center>
    </div>
  );
};

export default MyAppointmentAdd;
