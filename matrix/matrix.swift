import Foundation

class Matrix {
    
    var M : [[Int]]
    
    init(size: Int) {
        self.M = Matrix.generateMatrix(size)
    }
    
    class func generateMatrix(size: Int) -> [[Int]]{
        
        var Mret = [[Int]](count:size, repeatedValue: [Int](count: size, repeatedValue: 0))
        for y in 0..<size {
            
            for x in 0..<size {
                
                Mret[y][x] = y*size + x
            }
        }
        
        return Mret
    }
    
    func rotate1() {
        
        let n = M.count - 1
        var Mret = [[Int]](count:n+1, repeatedValue: [Int](count: n+1, repeatedValue: 0))
        
        for y in 0...n{
            
            for x in 0...n{
                
                Mret[x][n-y] = M[y][x]
            }
        }
        M = Mret
    }

    func rotate2() {
        
        let n = M.count - 1
        
        for y in 0...n {
            for x in 0...y {
                let tmp = M[x][y]
                M[x][y] = M[y][x]
                M[y][x] = tmp
            }
        }
        
        for y in 0...n {
            for x in 0...Int(floor(Double(n)/2)) {
                let tmp = M[y][n-x]
                M[y][n-x] = M[y][x]
                M[y][x] = tmp
            }
        }
        
    }

}

let size = 100
let matrix1 = Matrix(size:size)
matrix1.rotate1()

let matrix2 = Matrix(size:size)
matrix2.rotate2()

println(matrix1.M == matrix2.M)





