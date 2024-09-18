import { useLocation } from "react-router-dom";
import axios from "axios";
import { useMemo, useState } from "react";

interface Appointment {
  id: number;
  patientId: number;
  doctorId: number;
  dateTime: string;
  patientName: string;
}

interface Patient {
  fio: string;
  phoneNumber: string;
  email: string;
  insurance: string;
  id: number;
}

const AppointmentsGet: React.FC = () => {
  const location = useLocation();

  const [appointments, setAppointments] = useState<Appointment[]>([]);
  const [isLoading, setIsLoading] = useState<boolean>(false);

  const headers = useMemo(() => {
    return {
      "Content-Type": "application/json",
      "auth-x": `Bearer ${location.state.token}`,
    };
  }, [location.state.token]);

  const handleGetAppointmentByDoctorClick = async () => {
    setIsLoading(true);
    try {
      const response = await axios.get(
        `http://localhost:8080/api/v1/appointment?doctor_id=${location.state.did}`,
        { headers: headers },
      );
      const appointmentsData = response.data;

      const appointmentsWithPatients = await Promise.all(
        appointmentsData.map(async (appointment: Appointment) => {
          const responsePat = await axios.get<Patient>(
            `http://localhost:8080/api/v1/patient/${appointment.patientId}`,
            { headers: headers },
          );
          return { ...appointment, patientName: responsePat.data.fio };
        }),
      );

      setAppointments(appointmentsWithPatients);
    } catch (error) {
      alert("Ошибка получения записей");
      console.error("There was an error!", error);
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div>
      <center>
        <button
          style={{
            marginTop: window.innerHeight * 0.2,
            borderRadius: 10,
            backgroundColor: "cyan",
            height: 50,
            fontSize: 16,
          }}
          onClick={handleGetAppointmentByDoctorClick}
          disabled={isLoading}
        >
          {isLoading ? "Загрузка..." : "Получить мои записи"}
        </button>

        <h2>Ваши записи</h2>
        <table style={{ borderCollapse: "collapse", width: "100%" }}>
          <thead>
            <tr style={{ borderBottom: "1px solid #ddd" }}>
              <th style={{ textAlign: "center", padding: "8px" }}>
                ID пациента
              </th>
              <th style={{ textAlign: "center", padding: "8px" }}>
                ФИО пациента
              </th>
              <th style={{ textAlign: "center", padding: "8px" }}>
                Дата и время записи
              </th>
            </tr>
          </thead>
          <tbody>
            {appointments.map((item) => (
              <tr key={item.id} style={{ borderBottom: "1px solid #ddd" }}>
                <td style={{ textAlign: "center", padding: "8px" }}>
                  {item.patientId}
                </td>
                <td style={{ textAlign: "center", padding: "8px" }}>
                  {item.patientName}
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
            ))}
          </tbody>
        </table>
      </center>
    </div>
  );
};

export default AppointmentsGet;
