import psycopg2
from psycopg2 import sql


def get_table_data(cursor, table_name):
    cursor.execute(sql.SQL("SELECT * FROM {}")
                   .format(sql.Identifier(table_name)))
    return cursor.fetchall()


def main():
    conn = psycopg2.connect(
        dbname="dicdoc_service",
        user="postgres",
        password="postgres",
        host="localhost",
        port="5432"
    )
    cursor = conn.cursor()

    table_name = 'appointments'

    initial_data = get_table_data(cursor, table_name)

    doc_id = pat_id = 100
    datetime = "2024-06-11 19:00:00.000000"

    cursor.execute("SELECT * FROM add_appointment(%s, %s, %s)", (doc_id, pat_id, datetime,))
    conn.commit()

    final_data = get_table_data(cursor, table_name)

    initial_set = set(initial_data)
    final_set = set(final_data)

    print(sorted(list(initial_data)), sorted(list(final_data)), sep="\n\n")

    added_record = final_set - initial_set

    cursor.close()
    conn.close()

    print("\nДобавленная запись:", added_record)


if __name__ == "__main__":
    main()
