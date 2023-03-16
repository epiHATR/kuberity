class statusHelper {
    getReadyOnTotal(statuses) {
        const rawObject = JSON.parse(JSON.stringify(statuses));
        const totalPod = rawObject.length
        var ready = 0
        rawObject.forEach(item => {
            if (item.ready) {
                ready++
            }
        });
        return ready + "/" + totalPod
    }

    countRestart(containerStatuses) {
        const rawObject = JSON.parse(JSON.stringify(containerStatuses));
        var total = 0
        rawObject.forEach(element => {
            total += element.restartCount
        });
        return total
    }
}

export default new statusHelper();