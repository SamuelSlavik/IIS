export const formatDate = (isoDateString: string) => {
    const date = new Date(isoDateString);
    const day = date.getDate().toString().padStart(2, '0');
    const month = (date.getMonth() + 1).toString().padStart(2, '0');
    const year = date.getFullYear();
    return `${day}.${month}.${year}`;
}

export const formatDateTime = (isoDateString: string) => {
    const date = new Date(isoDateString);
    const day = date.getDate().toString().padStart(2, '0');
    const month = (date.getMonth() + 1).toString().padStart(2, '0');
    const year = date.getFullYear();
    const hour = date.getHours().toString().padStart(2, '0');
    const minute = date.getMinutes().toString().padStart(2, '0');
    return `${day}.${month}.${year} ${hour}:${minute}`;
}

export const formatTimeForCreate = (timeString: string): string => {
    const inputDate = new Date(timeString);
    if (isNaN(inputDate.getTime())) {
        return "";
    }
    const formattedTime = `${inputDate.toISOString().slice(0, 10)} ${inputDate.toISOString().slice(11, 16)}`;

    return formattedTime;
};
