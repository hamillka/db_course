import React, { useState } from "react";
import axios from "axios";
import { useLocation, useNavigate } from "react-router-dom";
import DatePicker from "react-datepicker";
import "react-datepicker/dist/react-datepicker.css";

const AppointmentEdit: React.FC = () => {
  const [id, setId] = useState(0);
  const [pid, setPid] = useState(0);
  const [did, setDid] = useState(0);
  const [date, setDate] = useState<Date | null>(null);

  const location = useLocation();
  const navigate = useNavigate();

  let headers = {
    "Content-Type": "application/json",
    "auth-x": `Bearer ${location.state.token}`,
  };

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

  const handleEditClick = async () => {
    try {
      const rfc3339 = dateToRFC3339WithOffset(date!, 3);
      await axios.patch(
        `http://localhost:8080/api/v1/appointment/${id}`,
        { patientId: pid, doctorId: did, dateTime: rfc3339 },
        { headers: headers },
      );
      navigate("/api/admin", {
        state: {
          login: location.state.login,
          password: location.state.password,
          token: location.state.token,
          uid: location.state.uid,
        },
      });
    } catch (error) {
      alert("Ошибка изменения записи");
      console.error("There was an error!", error);
    }
  };

  return (
    <div>
      <center>
        <input
          style={{ marginTop: window.innerHeight * 0.3, height: 30 }}
          type="number"
          placeholder="ID записи"
          onChange={(e) => setId(parseInt(e.target.value))}
        />
        <input
          style={{ height: 30 }}
          type="number"
          placeholder="ID пациента"
          onChange={(e) => setPid(parseInt(e.target.value))}
        />
        <input
          style={{ height: 30 }}
          type="number"
          placeholder="ID врача"
          onChange={(e) => setDid(parseInt(e.target.value))}
        />
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
          filterTime={(time) => {
            const selectedHour = time.getHours();
            return selectedHour >= 8 && selectedHour < 20;
          }}
        />
        <button
          style={{
            height: 30,
            borderRadius: 7.5,
            fontSize: 16,
            marginLeft: 7,
            marginTop: 15,
            width: 180,
            backgroundColor: "cyan",
          }}
          onClick={handleEditClick}
        >
          Изменить
        </button>
      </center>
    </div>
  );
};

export default AppointmentEdit;
