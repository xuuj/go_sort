package main

import (
	"fmt"
)

func main(){
	sortArr:=[]int{22,7,1,22,11,44,21,99,34,33,12,44,44,44,55}
	//quickSort(sortArr,0,len(sortArr)-1)
	//n:=time.Now().UnixNano()
	//bubbleSort(sortArr)
	//fmt.Println()
	//heapSort(sortArr)
	//selectSort(sortArr)
	//insertSort(sortArr)
	//mergeSort(sortArr,0,len(sortArr)-1)
	shellSort(sortArr)
	printArr(sortArr)

}

/**
快速排序
快速排序用一个数据做分界线，把其他数据放在这个数据的两边，在对分界线两边的数据做同样的操作直到有序
实现思路：
1.用第一个值作为基准base，i++,j--,找到第一个比base大的，和第一个比base小的，交换
2.到i=j 跳出，把base和i值交换，此时数组就已经变成左边比base小，右边比base大了
3.对左边，右边数组重复1,2（递归）。
快速排序的最好时间复杂度为O(logn)，最坏的时间复杂度为O(n^2)且其为不稳定排序。
 */
func quickSort(sortArr []int,start int,end int){
	if start>=end{
		return
	}
	i:=start
	j:=end
	base:=sortArr[start]
	for{
		for sortArr[j]>=base&&i<j{
			j--
		}

		for sortArr[i]<=base&&i<j{
			i++
		}

		if i>=j{
			break
		}

		sortArr[i],sortArr[j]=sortArr[j],sortArr[i]

	}
	sortArr[start],sortArr[i]=sortArr[i],sortArr[start]
	quickSort(sortArr,start,i-1)
	quickSort(sortArr,i+1,end)

}


func printArr(sortArr []int){
	for i:=0;i<len(sortArr);i++{
		fmt.Print(sortArr[i],",")

	}
	fmt.Println("")
}
/**
 数组为大小为n
 进行n趟排序，每次如果大于相邻的值则交换
 冒泡排序为稳定排序
 最好的时间复杂度为O(n)即排序数组开始就是有序的且与要排序的顺序相同，因此只要进行一趟比较排序，遍历一次数组。最坏的情况自然是数组开始时与要排序的顺序相反，所以时间复杂度为O(n^2)
 */
func bubbleSort(sortArr []int){
	for i:=0;i<len(sortArr);i++{
		for j:=i;j<len(sortArr);j++{
			if sortArr[i]>sortArr[j]{
				sortArr[i],sortArr[j]=sortArr[j],sortArr[i]
			}
		}
	}

}


/**
	堆排序
	arr数组，len 为 n，从左往右形成了二叉树
	1.先对每个非叶子结点进行调整（最后一个非叶子结点，n/2-1，所以从数组0～n/2-1之前都是非叶子节点）
		.调整，根据非叶子结点j,可以得到它的两个叶子节点为2j+1,2j+2.
	2.全部调整完取出根节点，和最后一个叶子节点进行交换
	3.重复1,2直到只剩一个值
 */


 func heapSort(arr []int){
	n:=len(arr)
	for {
		if n==0{
			break
		}
		for j:=n/2-1;j>=0;j--{
			minHeap(arr,j,n-1)
		}
		arr[0],arr[n-1]=arr[n-1],arr[0]
		n--
	}

 }

 func maxHeap(arr []int,topIdx int,tail int){
 	topNode:=arr[topIdx]
 	leftNodeIdx:=topIdx*2+1
 	rightNodeIdx:=topIdx*2+2
 	var leftNode,rightNode int
 	if leftNodeIdx<=tail{
 		leftNode=arr[leftNodeIdx]
	}
	if rightNodeIdx<=tail{
		rightNode=arr[rightNodeIdx]
	}
	if leftNode>=rightNode&&leftNode>topNode{
		arr[topIdx],arr[leftNodeIdx]=arr[leftNodeIdx],arr[topIdx]
	}else if rightNode>=leftNode&&rightNode>topNode{
		arr[topIdx],arr[rightNodeIdx]=arr[rightNodeIdx],arr[topIdx]
	}
 }

 func minHeap(arr []int,topIdx int,tail int){
	 topNode:=arr[topIdx]
	 leftNodeIdx:=topIdx*2+1
	 rightNodeIdx:=topIdx*2+2
	 var leftNode,rightNode int
	 if leftNodeIdx>tail{
	 	return
	 }
	 leftNode=arr[leftNodeIdx]

	 if rightNodeIdx>tail{
	 	if leftNode<topNode {
			arr[topIdx], arr[leftNodeIdx] = arr[leftNodeIdx], arr[topIdx]
		}
		return
	 }

	 rightNode=arr[rightNodeIdx]
	 if leftNode<=rightNode&&leftNode<topNode{
		 arr[topIdx],arr[leftNodeIdx]=arr[leftNodeIdx],arr[topIdx]
	 }else if rightNode<=leftNode&&rightNode<topNode{
		 arr[topIdx],arr[rightNodeIdx]=arr[rightNodeIdx],arr[topIdx]
	 }

 }


 /**
 直接选择排序
 把数组分为有序区和无序区，每次从无序区中选择一个值放到有序区末尾
 实现：
 1.从数组中选出最小，最大值
 2.与数组第i个值交换（i=0,i++），就是每次把排完序的值放到前端
 3.直到所有值都排序完

 直接选择排序为不稳定排序
 稳定排序：在排序结束后如果排序中相同大小的数与排序前的相对位置相同(即前后位置不变)，则称为稳定排序
  */

  func selectSort(arr []int){
  	for i:=0;i<len(arr);i++{
		maxIdx:=i
  		for j:=i;j<len(arr);j++{
  			if arr[j]>arr[maxIdx]{
  				maxIdx=j
			}
		}
		if maxIdx!=i{
			arr[i],arr[maxIdx]=arr[maxIdx],arr[i]
		}
	}
  }

  /**
   直接插入排序
   数组的长度为n
   1.从数组1开始为基准值，往前比较，如果比前值小就交换
   2.如果不比前值小则直接break
   3.到数组最后一个值排序完结束
  通过上面的思想可以看出该排序为稳定排序，且最好时间复杂度、最坏时间复杂度都与前面的冒泡排序一样是O(n)与O(n^2)
   */

   func insertSort(arr []int){
   		for i:=1;i<len(arr);i++{
   			baseIdx:=i
   			for j:=i-1;j>=0;j--{
   				if arr[baseIdx]<arr[j]{
   					arr[baseIdx],arr[j]=arr[j],arr[baseIdx]
   					baseIdx=j
				}else{
					break
				}
			}
		}
   }

   /**
		归并排序
   	利用分治的思想，
   1.把需要排序的数组划分成两个数组，在对这两个数组再划分
   2.划分到只有一个值然后在合并两个数组。利用临时数组进行有序合并
   3.递归等到全部合并完排序结束

   将序列划分成组需要logn次，将划分的组合并的时间复杂度为O(n)，所以最终的时间复杂度为O(nlogn)
    */

    func mergeArr(arr []int,low int,mid int,high int){
		tempArr:=make([]int,high-low+1,high-low+1)
		i:=0
		base1:=low
		base2:=mid+1
		for{
			if base1>mid&&base2>high{
				break
			}
			if base1<=mid&&base2<=high{
				if arr[base2]>=arr[base1]{
					tempArr[i]=arr[base2]
					base2++
				}else{
					tempArr[i]=arr[base1]
					base1++
				}

			}else if base1>mid{
				tempArr[i]=arr[base2]
				base2++
			}else if base2>high{
				tempArr[i]=arr[base1]
				base1++
			}
			i++

		}
		for k:=0;k<len(tempArr);k++{
			arr[low+k]=tempArr[k]
		}
	}

	func mergeSort(arr1 []int,low int,high int){

		if low<high{
			mid:=(low+high)/2
			mergeSort(arr1,low,mid)
			mergeSort(arr1,mid+1,high)
			mergeArr(arr1,low,mid,high)
		}
	}


	/**
	 希尔排序
	希尔排序，对插入排序的改善，
	1.选择一个增量d（一般为数组长度n/2），然后从数组0开始每次选择一个长度为d的数组进行插入排序。如：[0,d][1,d+1]直到[i,d+i] d+n等于数组长度=n
	2.再次选择一个增量d（一般为d/2）重复1。
	3.直到d=1结束排序
	平均时间复杂度：希尔排序的时间复杂度和其增量序列有关系，这涉及到数学上尚未解决的难题；不过在某些序列中复杂度可以为O(n1.3);
	空间复杂度：O(1)
	稳定性：不稳定
	 */

	 func shellSort(arr []int){
	 	d:=len(arr)/2
	 	for d>1{
	 		for i:=0;i+d<=len(arr);i++{
	 			insertSort(arr[i:i+d])
			}
	 		d=d/2
		}
	 }