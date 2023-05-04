
export const isZipFile = (file: File): Promise<unknown> => {
    const fileReader = new FileReader();

    const isValidZipFileType = (fileType: string): boolean => {
        const zipFileTypes = ['application/zip', 'application/x-zip', 'application/octet-stream', 'application/x-zip-compressed'];
        return zipFileTypes.includes(fileType);
    };

    const checkFileType = (): boolean => {
        const arr = new Uint8Array(fileReader.result as ArrayBuffer).subarray(0, 4);
        let header = '';
        for (let i = 0; i < arr.length; i++) {
            header += arr[i].toString(16);
        }

        switch (header) {
            case '504b0304':
            case '504b0506':
            case '504b0708':
                return isValidZipFileType('application/zip');
            case '504b030414':
                return isValidZipFileType('application/x-zip-compressed');
            case '504b0508':
                return isValidZipFileType('application/x-zip');
            case '504b5370':
                return isValidZipFileType('application/octet-stream');
            default:
                return false;
        }
    };

    return new Promise((resolve, reject) => {
        fileReader.onloadend = () => {
            const fileType = file.type;
            if (fileType === '' || fileType === 'application/octet-stream') {
                // 如果浏览器不能识别文件类型，则进行手动检查
                resolve(checkFileType());
            } else {
                // 如果浏览器可以识别文件类型，则根据文件类型进行检查
                resolve(isValidZipFileType(fileType));
            }
        };

        fileReader.readAsArrayBuffer(file.slice(0, 4));
    });
}
