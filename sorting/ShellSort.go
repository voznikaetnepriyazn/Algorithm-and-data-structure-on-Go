package algs

func shellSort(arr []int) {
	gap := len(arr) / 2
	for i := gap; i < len(arr); i++ {
		m_gap := i
		for m_gap >= gap && arr[m_gap] < arr[m_gap-gap] {
			arr[m_gap], arr[m_gap-gap] = arr[m_gap-gap], arr[m_gap]
			m_gap -= gap
		}
	}
	gap = gap / 2
}
